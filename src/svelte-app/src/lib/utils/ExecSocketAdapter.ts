import type * as xterm from 'xterm';

import { b64decode, b64encode } from './Base64Util.js';

export const HEARTBEAT_INTERVAL = 10000;
export const MAX_RETRIES = 10;

export class ExecSocketAdapter {
	terminal: xterm.Terminal;
	socket: WebSocket;
	heartbeatTimer: NodeJS.Timer | undefined;
	retries: number;

	constructor(terminal: xterm.Terminal, url: string) {
		this.terminal = terminal;
		this.socket = new WebSocket(url);
		this.retries = 0;

		if (!this.terminal) {
			throw new Error('Terminal is not defined.');
		}
		this.terminal.onData((data) => {
			this.handleData(data);
		});
		this.connect();
	}

	connect() {
		this.socket.onopen = () => {
			this.sendWsHandshake();
			this.sendTtySize();
			this.startHeartbeat();
			this.terminal.reset();
		};

		this.socket.onmessage = (e) => {
			const json = JSON.parse(e.data);

			if (json.stdout && json.stdout.data) {
				this.retries = 0; // reset retries on successful message
				this.terminal.write(b64decode(json.stdout.data));
			}
		};

		this.socket.onclose = () => {
			this.stopHeartbeat();
			if (this.retries < MAX_RETRIES) {
				this.retries++;
				setTimeout(() => {
					this.reconnect();
				}, 1000);
			} else {
				this.terminal.reset();
				this.terminal.write('Failed to connect to server');
			}
		};

		this.terminal.onResize(() => {
			if (this.socket.readyState != WebSocket.OPEN) {
				return;
			}
			this.sendTtySize();
		});
	}

	reconnect() {
		console.log('reconnecting...');
		this.socket = new WebSocket(this.socket.url);
		this.connect();
	}

	sendTtySize() {
		this.socket.send(
			JSON.stringify({
				tty_size: { width: this.terminal.cols, height: this.terminal.rows }
			})
		);
	}

	sendWsHandshake() {
		this.socket.send(JSON.stringify({ version: 1, auth_token: '' }));
	}

	startHeartbeat() {
		this.heartbeatTimer = setInterval(() => {
			this.socket.send(JSON.stringify({}));
		}, HEARTBEAT_INTERVAL);
	}

	stopHeartbeat() {
		clearInterval(this.heartbeatTimer);
	}

	handleData(data: string) {
		if (this.socket.readyState != WebSocket.OPEN) {
			return;
		}
		this.socket.send(JSON.stringify({ stdin: { data: b64encode(data) } }));
	}

	close() {
		this.socket.close();
	}
}

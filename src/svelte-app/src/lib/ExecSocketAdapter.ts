import type * as xterm from 'xterm';
import { b64encode, b64decode } from './Base64Util.js';

export const HEARTBEAT_INTERVAL = 10000;

export class ExecSocketAdapter {
	terminal: xterm.Terminal;
	socket: WebSocket;
	heartbeatTimer: NodeJS.Timer | undefined;

	constructor(terminal: xterm.Terminal, url: string) {
		this.terminal = terminal;
		this.socket = new WebSocket(url);

		this.socket.onopen = () => {
			this.sendWsHandshake();
			this.sendTtySize();
			this.startHeartbeat();

			this.terminal.clear();
			this.terminal.onData((data) => {
				this.handleData(data);
			});
		};

		this.socket.onmessage = (e) => {
			const json = JSON.parse(e.data);

			if (json.stdout && json.stdout.data) {
				this.terminal.write(b64decode(json.stdout.data));
			}
		};

		this.socket.onclose = () => {
			this.stopHeartbeat();
			this.terminal.writeln('');
			this.terminal.writeln('Connection closed.');
		};
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
		if (this.socket.readyState == WebSocket.CLOSED) {
			return;
		}
		this.socket.send(JSON.stringify({ stdin: { data: b64encode(data) } }));
	}
}

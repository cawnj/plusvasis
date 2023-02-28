<script lang="ts">
	import { onMount } from 'svelte';
	import 'xterm/css/xterm.css';

	import { ExecSocketAdapter } from './ExecSocketAdapter';

	let xterm: any;
	let fit: any;

	let terminal: any;
	let terminalElement: HTMLElement;
	let termFit: any;
	let execSocketAdapter: ExecSocketAdapter;

	function initalizeXterm() {
		terminal = new xterm.Terminal({
			fontFamily: 'monospace',
			fontWeight: '400'
		});
		termFit = new fit.FitAddon();
		terminal.loadAddon(termFit);
		terminal.open(terminalElement);
		termFit.fit();
	}
	export function connectTerm(url: string) {
		execSocketAdapter = new ExecSocketAdapter(terminal, url);
	}
	onMount(async () => {
		xterm = await import('xterm');
		fit = await import('xterm-addon-fit');
		initalizeXterm();
	});
	export function write(content: string) {
		terminal.write(content);
	}
</script>

<div id="terminal" bind:this={terminalElement} />

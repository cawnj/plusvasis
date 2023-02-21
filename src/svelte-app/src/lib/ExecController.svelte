<script lang="ts">
	import { onMount } from 'svelte';
	import 'xterm/css/xterm.css';
	import * as xterm from 'xterm';
	import * as fit from 'xterm-addon-fit';

	import { ExecSocketAdapter } from './ExecSocketAdapter';

	let terminal: xterm.Terminal;
	let terminalElement: HTMLElement;
	let termFit: fit.FitAddon;
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
		initalizeXterm();
	});
	export function write(content: string) {
		terminal.write(content);
	}
</script>

<div id="terminal" bind:this={terminalElement} />

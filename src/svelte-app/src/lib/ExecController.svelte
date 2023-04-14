<script lang="ts">
	import { onMount } from 'svelte';
	import 'xterm/css/xterm.css';
	import type { Terminal } from 'xterm';
	import type { FitAddon } from 'xterm-addon-fit';
	import { ExecSocketAdapter } from './ExecSocketAdapter';

	type xtermType = typeof import('xterm');
	type fitType = typeof import('xterm-addon-fit');
	let xterm: xtermType;
	let fit: fitType;

	let terminal: Terminal;
	let terminalElement: HTMLElement;
	let termFit: FitAddon;

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
		new ExecSocketAdapter(terminal, url);
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

<div id="terminal" data-testid="exec-controller" bind:this={terminalElement} />

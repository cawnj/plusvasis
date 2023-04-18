<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import 'xterm/css/xterm.css';
	import type { Terminal } from 'xterm';
	import type { FitAddon } from 'xterm-addon-fit';
	import { ExecSocketAdapter } from './ExecSocketAdapter';

	export let url: string;
	let width: number;
	let fontSize = 14;

	$: {
		fontSize = width < 768 ? 12 : 18;
	}

	type xtermType = typeof import('xterm');
	type fitType = typeof import('xterm-addon-fit');
	let xterm: xtermType;
	let fit: fitType;
	let execSocketAdapter: ExecSocketAdapter;

	let terminal: Terminal;
	let terminalElement: HTMLElement;
	let termFit: FitAddon;

	function initalizeXterm() {
		terminal = new xterm.Terminal({
			fontFamily: 'monospace',
			fontWeight: '400',
			fontSize: fontSize,
			cursorBlink: true,
			scrollback: 10000
		});
		termFit = new fit.FitAddon();
		terminal.loadAddon(termFit);
	}
	function connectTerm() {
		execSocketAdapter = new ExecSocketAdapter(terminal, url);
	}
	export function write(content: string) {
		terminal.write(content);
	}
	function postInit() {
		if (terminalElement && terminal) {
			terminal.open(terminalElement);
			termFit.fit();
			connectTerm();
		} else {
			setTimeout(() => postInit(), 100);
		}
	}
	onMount(async () => {
		xterm = await import('xterm');
		fit = await import('xterm-addon-fit');
		initalizeXterm();
		postInit();
	});
	onDestroy(() => {
		terminal.dispose();
		execSocketAdapter.close();
	});
</script>

<svelte:window bind:innerWidth={width} />
<div id="terminal" data-testid="exec-controller" bind:this={terminalElement} />

<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import 'xterm/css/xterm.css';
	import type { Terminal } from 'xterm';
	import type { FitAddon } from 'xterm-addon-fit';
	import { ExecSocketAdapter } from './ExecSocketAdapter';
	import { currJobStopped } from '../stores/nomadStore';

	let isStopped: boolean;
	currJobStopped.subscribe((value) => {
		isStopped = value;
	});

	export let wsUrl: string;
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
		execSocketAdapter = new ExecSocketAdapter(terminal, wsUrl);
	}
	export function write(content: string) {
		terminal.write(content);
	}
	function postInit() {
		if (terminalElement && terminal) {
			terminal.open(terminalElement);
			termFit.fit();
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
		if (terminal) terminal.dispose();
		if (execSocketAdapter) execSocketAdapter.close();
	});
	$: if (!isStopped && wsUrl && terminal) {
		connectTerm();
	}
</script>

<svelte:window bind:innerWidth={width} />
<div id="terminal" data-testid="exec-controller" bind:this={terminalElement} />

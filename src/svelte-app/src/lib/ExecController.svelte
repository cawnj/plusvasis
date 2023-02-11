<script lang="ts">
	import { onMount } from 'svelte';
	import 'xterm/css/xterm.css';
	import * as xterm from 'xterm';
	import * as fit from 'xterm-addon-fit';

	let terminalElement: HTMLElement;
	let terminalController: xterm.Terminal;
	let termFit: fit.FitAddon;
	$: {
		if (terminalController) {
			// ...
		}
	}
	function initalizeXterm() {
		terminalController = new xterm.Terminal({
			convertEol: true
		});
		termFit = new fit.FitAddon();
		terminalController.loadAddon(termFit);
		terminalController.open(terminalElement);
		termFit.fit();
		terminalController.onData((e) => {
			console.log(e);
		});
	}
	onMount(async () => {
		initalizeXterm();
	});
	export function write(content: string) {
		terminalController.write(content);
	}
</script>

<div id="terminal" bind:this={terminalElement} />

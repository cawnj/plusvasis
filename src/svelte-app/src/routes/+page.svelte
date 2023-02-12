<script lang="ts">
	import ExecController from '../lib/ExecController.svelte';
	import Nav from '../lib/NavBar.svelte';

	let execControllerComponent: ExecController;

	function start_container() {
		execControllerComponent.write('Starting container...');
	}
</script>

<Nav />
<h1 class="mb-4 text-4xl font-bold font-sans text-white">Continens</h1>
<button class="mb-4 btn btn-blue" on:click={start_container}>Start</button>
<ExecController bind:this={execControllerComponent} />

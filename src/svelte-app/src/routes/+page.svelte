<script lang="ts">
	import ExecController from '$lib/ExecController.svelte';
	import Nav from '$lib/NavBar.svelte';
	export let data;
	const { jobs } = data;

	let execControllerComponent: ExecController;

	function start_container() {
		execControllerComponent.write('Starting container...');
	}

	let clicked = false;
	function get_containers() {
		clicked = true;
	}
</script>

<Nav />
<h1 class="mb-4 text-4xl font-bold font-sans text-white">Continens</h1>
<button class="mb-4 btn btn-blue" on:click={start_container}>Start</button>
<button class="mb-4 btn btn-blue" on:click={get_containers}>Get Containers</button>
{#if clicked}
	{#each jobs as job}
		<ul>
			<button
				type="button"
				class="focus:outline-none text-white bg-purple-700 hover:bg-purple-800 focus:ring-4 focus:ring-purple-300 font-medium rounded-lg text-sm px-5 py-2.5 mb-2 dark:bg-purple-600 dark:hover:bg-purple-700 dark:focus:ring-purple-900"
				on:click={start_container}>{job.ID}</button
			>
		</ul>
	{/each}
{/if}
<ExecController bind:this={execControllerComponent} />

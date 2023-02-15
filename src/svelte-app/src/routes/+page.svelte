<script lang="ts">
	import ExecController from '$lib/ExecController.svelte';
	import Nav from '$lib/NavBar.svelte';
	export let data;
	const { jobs } = data;

	let execControllerComponent: ExecController;
	let getContainerClicked = false;

	function startContainer() {
		execControllerComponent.write('Starting container...');
	}

	function getContainers() {
		getContainerClicked = true;
	}

	async function fetchJobId(jobId: string) {
		const url = 'http://localhost:8080/job/' + jobId;
		const res = await fetch(url);

		if (res.ok) {
			execControllerComponent.write('Starting container ' + jobId);
		} else {
			execControllerComponent.write('Error starting container ' + jobId);
		}
	}
</script>

<Nav />
<h1 class="mb-4 text-4xl font-bold font-sans text-white">Continens</h1>
<button class="mb-4 btn btn-blue" on:click={startContainer}>Start</button>
<button class="mb-4 btn btn-blue" on:click={getContainers}>Get Containers</button>
{#if getContainerClicked}
	{#each jobs as job}
		<ul>
			<button type="button" class="btn-purple" on:click={() => fetchJobId(job.ID)}>{job.ID}</button>
		</ul>
	{/each}
{/if}
<ExecController bind:this={execControllerComponent} />

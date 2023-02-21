<script lang="ts">
	import Nav from '$lib/NavBar.svelte';
	import NomadController from '$lib/NomadController.svelte';

	let getContainerClicked: boolean;
	let containerRunning: boolean;
	let job: string;

	export let data;
	const { jobs } = data;

	let nomadControllerComponent: NomadController;
</script>

<Nav />
<h1 class="mb-4 text-4xl font-bold font-sans text-white">Continens</h1>
<button class="mb-4 btn btn-blue" on:click={nomadControllerComponent.getContainers}
	>Get Containers</button
>
{#if getContainerClicked}
	{#each jobs as job}
		<ul>
			<button
				type="button"
				class="btn-purple"
				on:click={() => nomadControllerComponent.fetchJobId(job.ID)}>{job.ID}</button
			>
		</ul>
	{/each}
{/if}
<!-- disabled for the moment, breaks things -->
<!-- {#if containerRunning}
	<button
		type="button"
		class="btn-red"
		on:click={() => nomadControllerComponent.fetchJobIdDelete(job)}>Stop Container {job}</button
	>
{/if} -->
<NomadController
	bind:this={nomadControllerComponent}
	bind:getContainerClicked
	bind:containerRunning
	bind:job
/>

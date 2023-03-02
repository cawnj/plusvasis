<script lang="ts">
	import Nav from '$lib/NavBar.svelte';
	import NomadController from '$lib/NomadController.svelte';

	let getContainerClicked: boolean;
	let getContainerCreatedClicked: boolean;
	let containerRunning: boolean;
	let containerName = '';
	let dockerImage = '';

	export let data;
	const { jobs } = data;

	let nomadControllerComponent: NomadController;
</script>

<Nav />
<h1 class="mb-4 text-4xl font-bold font-sans text-white">Continens</h1>
<button class="mb-4 btn btn-blue" on:click={nomadControllerComponent.getContainers}
	>Get Containers</button
>
<button class="mb-4 btn btn-blue" on:click={nomadControllerComponent.getContainerCreated}
	>Create Container</button
>
{#if getContainerClicked}
	{#each jobs as job}
		<ul>
			<button
				type="button"
				class="btn-purple"
				on:click={() => nomadControllerComponent.fetchJobIdAllocations(job.ID)}>{job.ID}</button
			>
		</ul>
	{/each}
{/if}
{#if getContainerCreatedClicked}
	<div class="mb-3 mt-3">
		<label for="containerNameInput" class="txt-input-label">Container Name</label>
		<input
			type="containerName"
			class="txt-input"
			id="containerNameInput"
			aria-describedby="containerNameHelp"
			placeholder="Container Name"
		/>
	</div>
	<div class="mb-3">
		<label for="imageInput" class="txt-input-label">Docker Image</label>
		<input
			type="dockerImage"
			class="txt-input"
			id="dockerImageInput"
			aria-describedby="dockerImageNameHelp"
			placeholder="Docker Image"
		/>
	</div>
	<div class="mb-3">
		<label for="imageInput" class="txt-input-label">Ports</label>
		<input
			type="ports"
			class="txt-input"
			id="portsInput"
			aria-describedby="portsHelp"
			placeholder="Ports"
		/>
	</div>
	<button class="mb-4 btn btn-blue" on:click={nomadControllerComponent.fetchJobCreate}
		>Create Container</button
	>
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
	bind:getContainerCreatedClicked
	bind:containerRunning
/>

<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import Nav from '$lib/NavBar.svelte';
	import NomadController from '$lib/NomadController.svelte';

	let nomadControllerComponent: NomadController;

	let containerName: '';
	let dockerImage: '';
	let validPath = false;

	async function fetchJobId() {
		const jobId = $page.params.id;
		const url = `http://localhost:8080/job/${jobId}`;
		const res = await fetch(url);
		const data = await res.json();

		if (res.ok) {
			return data;
		} else {
			throw new Error(data);
		}
	}

	fetchJobId().then(
		(job) => {
			validPath = true;
			containerName = job.ID;
			dockerImage = job.TaskGroups[0].Tasks[0].Config.image;
		},
		(err) => {
			console.error('Could not reach backend', err);
			validPath = false;
		}
	);
</script>

{#if validPath}
	<Nav />
	<h1 class="mb-4 text-4xl font-bold font-sans text-white">{$page.params.id}</h1>
	<div class="mb-3 mt-3">
		<label for="containerNameInput" class="txt-input-label">Container Name</label>
		<input
			type="containerName"
			class="txt-input"
			id="containerNameInput"
			aria-describedby="containerNameHelp"
			placeholder="Container Name"
			value={containerName}
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
			value={dockerImage}
		/>
	</div>
	<button
		class="mb-4 btn btn-blue"
		on:click={() => nomadControllerComponent.fetchJobUpdate(containerName)}>Update Container</button
	>
	<div hidden>
		<NomadController bind:this={nomadControllerComponent} />
	</div>
{/if}
{#if !validPath}
	<h1 class="mb-4 text-4xl font-bold font-sans text-white">Page Not Found</h1>
	<button class="mb-4 btn btn-blue" on:click={() => goto('/')}>Return to Homepage</button>
{/if}

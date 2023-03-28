<script lang="ts">
	import { goto } from '$app/navigation';
	import NavBar from '$lib/NavBar.svelte';
	import type { Job } from '$lib/Types';
	import { hostname } from '../../stores/environmentStore';

	async function fetchJobCreate(job: Job) {
		const url = `${hostname}/jobs`;
		const res = await fetch(url, {
			method: 'POST',
			body: JSON.stringify(job),
			headers: {
				Authorization: `Bearer ${localStorage.getItem('token')}`
			}
		});

		if (res.ok) {
			console.log('Container Created');
		} else {
			console.log('Error');
		}
		goto('/');
	}

	async function createJob() {
		const containerName = document.getElementById('containerNameInput') as HTMLInputElement;
		const dockerImage = document.getElementById('dockerImageInput') as HTMLInputElement;
		const shell = document.getElementById('shellInput') as HTMLInputElement;

		const job: Job = {
			containerName: containerName.value,
			dockerImage: dockerImage.value,
			user: localStorage.getItem('uid'),
			shell: shell.value
		};
		fetchJobCreate(job);
	}
</script>

<NavBar />
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
		aria-describedby="dockerImageHelp"
		placeholder="Docker Image"
	/>
</div>
<div class="mb-3">
	<label for="shellInput" class="txt-input-label">Shell Command</label>
	<input
		type="shell"
		class="txt-input"
		id="shellInput"
		aria-describedby="shellNameHelp"
		placeholder="Shell Command"
	/>
</div>
<button class="mb-4 btn btn-blue" on:click={() => createJob()}>Create Container</button>

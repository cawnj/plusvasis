<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import Nav from '$lib/NavBar.svelte';
	import type { Job } from '$lib/Types';
	import { job } from '../../../../stores/nomadStore';
	import { onMount } from 'svelte';
	import { hostname } from '../../../../stores/environmentStore';

	let newJob: Job = {
		containerName: '',
		dockerImage: '',
		user: ''
	};
	let jobId: string;
	job.subscribe((value) => {
		jobId = value;
	});
	let jobName: string;
	onMount(async () => {
		const jobId = $page.params.id;
		job.set(jobId);

		const url = `${hostname}/job/${jobId}`;
		const res = await fetch(url, {
			headers: {
				Authorization: `Bearer ${localStorage.getItem('token')}`
			}
		});
		if (res.ok) {
			const data = await res.json();
			jobName = data.Name;
			newJob.containerName = data.Name;
			newJob.dockerImage = data.TaskGroups[0].Tasks[0].Config.image;
		}
	});

	async function fetchJobUpdate(job: Job) {
		const url = `${hostname}/job/${jobId}`;
		const res = await fetch(url, {
			method: 'POST',
			body: JSON.stringify(job),
			headers: {
				Authorization: `Bearer ${localStorage.getItem('token')}`
			}
		});

		if (res.ok) {
			console.log('Container Updated');
		} else {
			console.log('Error');
		}
		goto('/');
	}

	async function updateJob() {
		const dockerImage = document.getElementById('dockerImageInput') as HTMLInputElement;

		newJob.dockerImage = dockerImage.value;
		newJob.user = localStorage.getItem('uid');
		fetchJobUpdate(newJob);
	}
</script>

<Nav />
{#if jobName}
	<h1 class="mb-4 text-4xl font-bold font-sans text-white">{jobName}</h1>
	<div class="mb-3">
		<label for="imageInput" class="txt-input-label">Docker Image</label>
		<input
			type="dockerImage"
			class="txt-input"
			id="dockerImageInput"
			aria-describedby="dockerImageNameHelp"
			placeholder="Docker Image"
			value={newJob.dockerImage}
		/>
	</div>
	<button class="mb-4 btn btn-blue" on:click={() => updateJob()}>Update Container</button>
{:else}
	<h1 class="mb-4 text-4xl font-bold font-sans text-white">Page Not Found</h1>
	<button class="mb-4 btn btn-blue" on:click={() => goto('/')}>Return to Homepage</button>
{/if}

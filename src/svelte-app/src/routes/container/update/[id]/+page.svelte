<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import Nav from '$lib/NavBar.svelte';
	import type { Job } from '$lib/Types';
	import { job } from '../../../../stores/nomadStore';
	import { onMount } from 'svelte';
	import { hostname } from '../../../../stores/environmentStore';

	let oldJob = {} as Job;
	let newJob = {} as Job;
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
			oldJob.containerName = data.Name;
			oldJob.dockerImage = data.TaskGroups[0].Tasks[0].Config.image;
			oldJob.shell = data.Meta.shell;
			oldJob.volumes = data.Meta.volumes;
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
		const shell = document.getElementById('shellInput') as HTMLInputElement;
		const volumeStr = document.getElementById('volumesInput') as HTMLInputElement;

		const volumes: [string, string][] = [];
		for (const volume of volumeStr.value.split(',')) {
			if (volume === '') {
				continue;
			}
			volumes.push(volume.split(':') as [string, string]);
		}

		newJob.containerName = oldJob.containerName;
		newJob.dockerImage = dockerImage.value;
		newJob.user = localStorage.getItem('uid');
		newJob.shell = shell.value;
		newJob.volumes = volumes;
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
			aria-describedby="dockerImageHelp"
			placeholder="Docker Image"
			value={oldJob.dockerImage}
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
			value={oldJob.shell}
		/>
	</div>
	<div class="mb-3">
		<label for="volumesInput" class="txt-input-label">Volumes</label>
		<input
			type="volumes"
			class="txt-input"
			id="volumesInput"
			aria-describedby="volumesHelp"
			placeholder="Volumes"
			value={oldJob.volumes}
		/>
	</div>
	<button class="mb-4 btn btn-blue" on:click={() => updateJob()}>Update Container</button>
{:else}
	<h1 class="mb-4 text-4xl font-bold font-sans text-white">Page Not Found</h1>
	<button class="mb-4 btn btn-blue" on:click={() => goto('/')}>Return to Homepage</button>
{/if}

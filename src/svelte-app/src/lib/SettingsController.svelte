<script lang="ts">
	import { JobFields, type Job } from '$lib/Types';
	import { currJobId, currJob } from '../stores/nomadStore';
	import { onMount } from 'svelte';
	import { hostname } from '../stores/environmentStore';

	let jobId: string;
	let job: Job;
	currJobId.subscribe((value) => {
		jobId = value;
	});
	currJob.subscribe((value) => {
		job = value;
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
	}

	async function updateJob() {
		const dockerImage = document.getElementById('dockerImageInput') as HTMLInputElement;
		const shell = document.getElementById('shellInput') as HTMLInputElement;
		const volumeStr = document.getElementById('volumesInput') as HTMLInputElement;
		const envStr = document.getElementById('envInput') as HTMLInputElement;
		const port = document.getElementById('portInput') as HTMLInputElement;

		const volumes: [string, string][] = [];
		for (const volume of volumeStr.value.split(',')) {
			if (volume === '') {
				continue;
			}
			volumes.push(volume.split(':') as [string, string]);
		}
		const envs: [string, string][] = [];
		for (const env of envStr.value.split(',')) {
			if (env === '') {
				continue;
			}
			envs.push(env.split('=') as [string, string]);
		}

		const newJob: Job = {
			user: localStorage.getItem('uid'),
			containerName: job.containerName,
			dockerImage: dockerImage.value,
			shell: shell.value,
			volumes: volumes,
			env: envs,
			port: Number(port.value),
			expose: false
		};
		fetchJobUpdate(newJob);
		window.location.reload();
	}
</script>

{#each JobFields as { key, value }}
	{#if key !== 'containerName'}
		<div class="mb-3 mt-3">
			<label for="{key}Input" class="txt-input-label">{value.title}</label>
			<input
				type={key}
				class="txt-input"
				id="{key}Input"
				aria-describedby="{key}Help"
				placeholder={value.placeholder}
				value={job[key]}
			/>
			<p class="text-sm text-gray-400">{value.info}</p>
		</div>
	{/if}
{/each}
<button class="mb-4 btn btn-blue" on:click={() => updateJob()}>Update Container</button>

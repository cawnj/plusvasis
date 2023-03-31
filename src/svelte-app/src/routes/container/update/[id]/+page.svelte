<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import Nav from '$lib/NavBar.svelte';
	import { JobFields, type Job } from '$lib/Types';
	import { job } from '../../../../stores/nomadStore';
	import { onMount } from 'svelte';
	import { hostname } from '../../../../stores/environmentStore';

	let jobId: string;
	job.subscribe((value) => {
		jobId = value;
	});

	let oldJob = new Map<string, string>();
	let newJob = {} as Job;

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
			oldJob.set('containerName', data.Name);
			oldJob.set('dockerImage', data.TaskGroups[0].Tasks[0].Config.image);
			oldJob.set('shell', data.Meta.shell);
			oldJob.set('volumes', data.Meta.volumes);
			oldJob.set('env', data.Meta.env);
			oldJob.set('port', data.Meta.port);
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

		newJob.containerName = jobName;
		newJob.dockerImage = dockerImage.value;
		newJob.user = localStorage.getItem('uid');
		newJob.shell = shell.value;
		newJob.volumes = volumes;
		newJob.env = envs;
		newJob.port = Number(port.value);
		fetchJobUpdate(newJob);
	}
</script>

<Nav />
{#if jobName}
	<h1 class="mb-4 text-4xl font-bold font-sans text-white">{jobName}</h1>
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
					value={oldJob.get(key)}
				/>
				<p class="text-sm text-gray-400">{value.info}</p>
			</div>
		{/if}
	{/each}
	<button class="mb-4 btn btn-blue" on:click={() => updateJob()}>Update Container</button>
{:else}
	<h1 class="mb-4 text-4xl font-bold font-sans text-white">Page Not Found</h1>
	<button class="mb-4 btn btn-blue" on:click={() => goto('/')}>Return to Homepage</button>
{/if}

<script lang="ts">
	import { goto } from '$app/navigation';
	import NavBar from '$lib/NavBar.svelte';
	import { JobFields, type Job } from '$lib/Types';
	import { hostname } from '../../stores/environmentStore';

	let job = {} as Job;

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

		job.containerName = containerName.value;
		job.dockerImage = dockerImage.value;
		job.user = localStorage.getItem('uid');
		job.shell = shell.value;
		job.volumes = volumes;
		job.env = envs;
		job.port = Number(port.value);
		fetchJobCreate(job);
	}
</script>

<NavBar />
{#each JobFields as { key, value }}
	<div class="mb-3 mt-3">
		<label for="{key}Input" class="txt-input-label">{value.title}</label>
		<input
			type={key}
			class="txt-input"
			id="{key}Input"
			aria-describedby="{key}Help"
			placeholder={value.placeholder}
		/>
		<p class="text-sm text-gray-400">{value.info}</p>
	</div>
{/each}
<button class="mb-4 btn btn-blue" on:click={() => createJob()}>Create Container</button>

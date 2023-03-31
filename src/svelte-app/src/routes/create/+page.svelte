<script lang="ts">
	import { goto } from '$app/navigation';
	import NavBar from '$lib/NavBar.svelte';
	import type { Job } from '$lib/Types';
	import { hostname } from '../../stores/environmentStore';
	import { Button } from 'flowbite-svelte';
	import JobForm from '$lib/JobForm.svelte';

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

	const createJob = () => {
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

		const job: Job = {
			user: localStorage.getItem('uid'),
			containerName: containerName.value,
			dockerImage: dockerImage.value,
			shell: shell.value,
			volumes: volumes,
			env: envs,
			port: Number(port.value),
			expose: false
		};
		fetchJobCreate(job);
	};
</script>

<NavBar />
<JobForm />
<Button color="blue" on:click={createJob}>Create Container</Button>

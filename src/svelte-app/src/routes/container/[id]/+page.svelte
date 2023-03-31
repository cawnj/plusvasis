<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import Nav from '$lib/NavBar.svelte';
	import NomadController from '$lib/NomadController.svelte';
	import Tabs from '$lib/Tabs.svelte';
	import type { Job, Tab } from '$lib/Types';
	import { currJobId, currJob } from '../../../stores/nomadStore';
	import { onMount } from 'svelte';
	import { hostname } from '../../../stores/environmentStore';
	import LogController from '$lib/LogController.svelte';
	import SettingsController from '$lib/SettingsController.svelte';
	import { Button } from 'flowbite-svelte';

	let jobName: string;
	onMount(async () => {
		const jobId = $page.params.id;
		currJobId.set(jobId);

		const url = `${hostname}/job/${jobId}`;
		const res = await fetch(url, {
			headers: {
				Authorization: `Bearer ${localStorage.getItem('token')}`
			}
		});
		if (res.ok) {
			const data = await res.json();
			jobName = data.Name;

			const job: Job = {
				user: localStorage.getItem('uid'),
				containerName: data.Name,
				dockerImage: data.TaskGroups[0].Tasks[0].Config.image,
				shell: data.Meta.shell,
				volumes: data.Meta.volumes,
				env: data.Meta.env,
				port: data.Meta.port,
				expose: false
			};
			currJob.set(job);
		}
	});

	const tabs: Tab[] = [
		{
			name: 'Shell',
			component: NomadController
		},
		{
			name: 'Logs',
			component: LogController
		},
		{
			name: 'Settings',
			component: SettingsController
		}
	];
</script>

{#if jobName}
	<Nav />
	<h1 class="mb-4 text-4xl font-bold font-sans text-white">{jobName}</h1>
	<Tabs {tabs} />
{:else}
	<h1 class="mb-4 text-4xl font-bold font-sans text-white">Page Not Found</h1>
	<Button color="blue" href="/">Return to Homepage</Button>
{/if}

<script lang="ts">
	import { page } from '$app/stores';
	import Nav from '$lib/NavBar.svelte';
	import NomadController from '$lib/NomadController.svelte';
	import Tabs from '$lib/Tabs.svelte';
	import type { Tab } from '$lib/Types';
	import { currJobId, currJob } from '../../../stores/nomadStore';
	import { onMount } from 'svelte';
	import LogController from '$lib/LogController.svelte';
	import SettingsController from '$lib/SettingsController.svelte';
	import { Button } from 'flowbite-svelte';
	import { faTerminal, faFileAlt, faCog } from '@fortawesome/free-solid-svg-icons';
	import { fetchJob } from '$lib/NomadClient';

	let jobName: string;
	onMount(async () => {
		const jobId = $page.params.id;
		currJobId.set(jobId);
		const job = await fetchJob(jobId);
		console.log(job);
		if (!job) {
			return;
		}
		jobName = job.containerName as string;
		currJob.set(job);
	});

	const tabs: Tab[] = [
		{
			name: 'Shell',
			component: NomadController,
			icon: faTerminal
		},
		{
			name: 'Logs',
			component: LogController,
			icon: faFileAlt
		},
		{
			name: 'Settings',
			component: SettingsController,
			icon: faCog
		}
	];
</script>

<Nav />
<div class="px-8 md:px-16">
	{#if jobName}
		<h1 class="mb-4 text-4xl font-bold font-sans text-white">{jobName}</h1>
		<Tabs {tabs} />
	{:else}
		<h1 class="mb-4 text-4xl font-bold font-sans text-white">Page Not Found</h1>
		<Button color="blue" href="/">Return to Homepage</Button>
	{/if}
</div>

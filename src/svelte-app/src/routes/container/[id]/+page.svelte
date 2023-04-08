<script lang="ts">
	import { page } from '$app/stores';
	import Nav from '$lib/NavBar.svelte';
	import NomadController from '$lib/NomadController.svelte';
	import Tabs from '$lib/Tabs.svelte';
	import type { Tab } from '$lib/Types';
	import { currJobId, currJob } from '../../../stores/nomadStore';
	import LogController from '$lib/LogController.svelte';
	import SettingsController from '$lib/SettingsController.svelte';
	import { Button, Modal, Spinner } from 'flowbite-svelte';
	import { faTerminal, faFileAlt, faCog } from '@fortawesome/free-solid-svg-icons';
	import { fetchJob } from '$lib/NomadClient';

	const fetchAndSetJob = async () => {
		const job = await fetchJob($page.params.id);
		if (!job) throw new Error('Job not found');
		currJobId.set($page.params.id);
		currJob.set(job);
		return job;
	};

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
<div class="px-8 md:px-16 mb-4">
	{#await fetchAndSetJob()}
		<div class="grid h-96 place-items-center">
			<Spinner />
		</div>
	{:then job}
		<h1 class="mb-4 text-4xl font-bold font-sans text-white">{job.containerName}</h1>
		<Tabs {tabs} />
	{:catch error}
		<Modal title="Error" open={true}>
			<div class="grid justify-center w-40">
				<p class="text-base leading-relaxed text-gray-500 dark:text-gray-400">{error.message}</p>
				<div class="flex justify-center mt-4">
					<Button color="blue" href="/">Return</Button>
				</div>
			</div>
		</Modal>
	{/await}
</div>

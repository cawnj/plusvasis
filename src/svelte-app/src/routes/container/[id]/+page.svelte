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
	import { faTerminal, faFileAlt, faCog, faExternalLink } from '@fortawesome/free-solid-svg-icons';
	import { fetchJob } from '$lib/NomadClient';
	import Fa from 'svelte-fa';
	import ContainerOptions from '$lib/ContainerOptions.svelte';

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
		<div class="flex justify-between items-center">
			<div class="flex items-center">
				<h1 class="mb-4 text-4xl font-bold font-sans text-white align-middle">
					{job.containerName}
				</h1>
				{#if job.port}
					<a class="align-middle mb-2" href="https://{$currJobId}.local.plusvasis.xyz">
						<Fa icon={faExternalLink} color="white" class="ml-2" />
					</a>
				{/if}
			</div>
			<ContainerOptions />
		</div>
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

<script lang="ts">
	import { page } from '$app/stores';
	import Nav from '$lib/NavBar.svelte';
	import NomadController from '$lib/NomadController.svelte';
	import Tabs from '$lib/Tabs.svelte';
	import type { Tab } from '$lib/Types';
	import { currJobId, currJob } from '../../../stores/nomadStore';
	import LogController from '$lib/LogController.svelte';
	import SettingsController from '$lib/SettingsController.svelte';
	import { Button, ButtonGroup, Modal, Spinner } from 'flowbite-svelte';
	import {
		faTerminal,
		faFileAlt,
		faCog,
		faPlay,
		faStop,
		faTrash,
		faRefresh
	} from '@fortawesome/free-solid-svg-icons';
	import { fetchJob } from '$lib/NomadClient';
	import Fa from 'svelte-fa';

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
			<h1 class="mb-4 text-4xl font-bold font-sans text-white">{job.containerName}</h1>
			<ButtonGroup>
				<Button>
					<Fa icon={faPlay} color="green" class="mr-2" />
					Start
				</Button>
				<Button>
					<Fa icon={faRefresh} color="orange" class="mr-2" />
					Restart
				</Button>
				<Button>
					<Fa icon={faStop} color="red" class="mr-2" />
					Stop
				</Button>
				<Button>
					<Fa icon={faTrash} class="mr-2" />
					Delete
				</Button>
			</ButtonGroup>
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

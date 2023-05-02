<script lang="ts">
	import { page } from '$app/stores';
	import NomadController from '$lib/NomadController.svelte';
	import Tabs from '$lib/Tabs.svelte';
	import type { Tab } from '$lib/Types';
	import LogController from '$lib/LogController.svelte';
	import SettingsController from '$lib/SettingsController.svelte';
	import { faTerminal, faFileAlt, faCog, faExternalLink } from '@fortawesome/free-solid-svg-icons';
	import Fa from 'svelte-fa';
	import ContainerOptions from '$lib/ContainerOptions.svelte';
	import { Heading, A, Modal, Button } from 'flowbite-svelte';
	import { onMount } from 'svelte';
	import { currJob, currJobId, currJobStopped } from '../stores/nomadStore';

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

	onMount(() => {
		currJobId.set($page.params.id);
		currJob.set($page.data.job);
		currJobStopped.set($page.data.isStopped);
	});
</script>

<div class="px-4 md:px-16 mb-4">
	{#if $page.data.job}
		<div class="flex justify-between mb-4">
			<div class="flex items-center font-sans font-bold">
				<Heading tag="h1" customSize="text-2xl md:text-4xl">
					{$page.data.job.containerName}
				</Heading>
				{#if $page.data.job.expose}
					<A aClass="ml-2" href="https://{$page.params.id}.plusvasis.xyz">
						<Fa icon={faExternalLink} color="white" size="xs" />
					</A>
				{/if}
			</div>
			<ContainerOptions />
		</div>
		<Tabs {tabs} />
	{:else}
		<Modal title="Error" open={true}>
			<div class="grid justify-center w-40">
				<p class="text-base leading-relaxed text-gray-500 dark:text-gray-400">{$page.data.error}</p>
				<div class="flex justify-center mt-4">
					<Button color="blue" href="/">Return</Button>
				</div>
			</div>
		</Modal>
	{/if}
</div>

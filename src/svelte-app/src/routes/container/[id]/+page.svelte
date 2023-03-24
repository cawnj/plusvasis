<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import Nav from '$lib/NavBar.svelte';
	import NomadController from '$lib/NomadController.svelte';
	import Tabs from '$lib/Tabs.svelte';
	import type { Tab } from '$lib/Types';
	import { job } from '../../../stores/nomadStore';
	import { onMount } from 'svelte';
	import { hostname } from '../../../stores/environmentStore';
	import LogController from '$lib/LogController.svelte';
	import SettingsController from '$lib/SettingsController.svelte';

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
	<div class="mb-2">
		<button
			type="button"
			class="btn-purple"
			on:click={() => goto('/container/update/' + $page.params.id)}>Update Container</button
		>
	</div>
	<Tabs {tabs} />
{:else}
	<h1 class="mb-4 text-4xl font-bold font-sans text-white">Page Not Found</h1>
	<button class="mb-4 btn btn-blue" on:click={() => goto('/')}>Return to Homepage</button>
{/if}

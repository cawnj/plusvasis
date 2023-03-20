<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import Nav from '$lib/NavBar.svelte';
	import NomadController from '$lib/NomadController.svelte';
	import { hostname } from '../../../stores/environmentStore';

	let nomadControllerComponent: NomadController;
	let validPath = false;
	let jobName: string;

	async function fetchJobId() {
		const jobId = $page.params.id;
		const url = `${hostname}/job/${jobId}`;
		const res = await fetch(url);
		const data = await res.json();

		if (res.ok) {
			return data;
		} else {
			throw new Error(data);
		}
	}

	fetchJobId().then(
		(data) => {
			if (data.Meta && data.Meta.user == localStorage.getItem('uid')) {
				validPath = true;
				jobName = data.Name;
			} else {
				validPath = false;
			}
		},
		(err) => {
			validPath = false;
		}
	);
</script>

{#if validPath}
	<Nav />
	<h1 class="mb-4 text-4xl font-bold font-sans text-white">{jobName}</h1>
	<div class="mb-2">
		<button
			type="button"
			class="btn-purple"
			on:click={() => nomadControllerComponent.fetchJobIdAllocations($page.params.id)}
			>Start Container</button
		>
		<button
			type="button"
			class="btn-purple"
			on:click={() => goto('/container/update/' + $page.params.id)}>Update Container</button
		>
	</div>
	<NomadController bind:this={nomadControllerComponent} />
{/if}
{#if !validPath}
	<h1 class="mb-4 text-4xl font-bold font-sans text-white">Page Not Found</h1>
	<button class="mb-4 btn btn-blue" on:click={() => goto('/')}>Return to Homepage</button>
{/if}

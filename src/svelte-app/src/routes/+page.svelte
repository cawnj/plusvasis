<script lang="ts">
	import { goto } from '$app/navigation';
	import Nav from '$lib/NavBar.svelte';
	import { onMount } from 'svelte';
	import { hostname } from '../stores/environmentStore';
	import Fa from 'svelte-fa';
	import { faTerminal } from '@fortawesome/free-solid-svg-icons';

	let jobs: any[] = [];
	onMount(async () => {
		const res = await fetch(`${hostname}/jobs`, {
			headers: {
				Authorization: `Bearer ${localStorage.getItem('token')}`
			}
		});
		if (res.ok) {
			jobs = await res.json();
		}
	});
</script>

<Nav />
<div class="">
	<button class="mb-4 btn btn-blue" on:click={() => goto('/create')}>Create Container</button>
	{#each jobs as job}
		<a href="/container/{job.ID}">
			<div class="div-container mt-3">
				<div class="flex items-center">
					<Fa icon={faTerminal} color="white" class="pr-6" />
					<span class="text-xl text-white">{job.Name}</span>
				</div>
			</div>
		</a>
	{/each}
</div>

<script lang="ts">
	import Nav from '$lib/NavBar.svelte';
	import { onMount } from 'svelte';
	import { hostname } from '../stores/environmentStore';
	import Fa from 'svelte-fa';
	import { faTerminal } from '@fortawesome/free-solid-svg-icons';
	import { logout } from '$lib/fb';
	import { Button } from 'flowbite-svelte';
	import { goto } from '$app/navigation';

	let jobs: object[] = [];
	onMount(async () => {
		const res = await fetch(`${hostname}/jobs`, {
			headers: {
				Authorization: `Bearer ${localStorage.getItem('token')}`
			}
		});
		if (res.ok) {
			jobs = await res.json();
		} else if (res.status === 401) {
			logout();
		}
	});
</script>

<Nav />
<div class="px-8 md:px-16">
	<Button color="blue" href="/create">Create Container</Button>
	{#if jobs}
		{#each jobs as job}
			<div
				class="div-container my-3 cursor-pointer"
				on:click={() => goto(`/container/${job.ID}`)}
				on:keydown={(event) => {
					if (event.key === 'Enter' || event.key === ' ') {
						goto(`/container/${job.ID}`);
					}
				}}
			>
				<button>
					<div class="flex items-center">
						<Fa icon={faTerminal} color="white" class="pr-6" />
						<span class="text-xl text-white">{job.Name}</span>
					</div>
				</button>
			</div>
		{/each}
	{/if}
</div>

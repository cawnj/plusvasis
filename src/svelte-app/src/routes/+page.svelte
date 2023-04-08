<script lang="ts">
	import Nav from '$lib/NavBar.svelte';
	import { hostname } from '../stores/environmentStore';
	import Fa from 'svelte-fa';
	import { faTerminal } from '@fortawesome/free-solid-svg-icons';
	import { logout } from '$lib/fb';
	import { Button, Spinner, Modal } from 'flowbite-svelte';
	import { goto } from '$app/navigation';

	const fetchJobs = async () => {
		let res: Response;
		try {
			res = await fetch(`${hostname}/jobs`, {
				headers: {
					Authorization: `Bearer ${localStorage.getItem('token')}`
				}
			});
		} catch (error) {
			throw new Error('Failed to fetch jobs');
		}

		if (res.ok) {
			return await res.json();
		} else if (res.status === 401) {
			logout();
		} else {
			throw new Error('Failed to fetch jobs');
		}
	};
</script>

<Nav />
<div class="px-8 md:px-16">
	{#await fetchJobs()}
		<div class="grid h-96 place-items-center">
			<Spinner />
		</div>
	{:then jobs}
		<Button color="blue" href="/create">Create Container</Button>
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
	{:catch error}
		<Modal title="Error" open={true}>
			<div class="grid justify-center w-40">
				<p class="text-base leading-relaxed text-gray-500 dark:text-gray-400">{error.message}</p>
				<div class="flex justify-center mt-4">
					<Button color="blue" href="javascript:window.location.href=window.location.href"
						>Retry</Button
					>
				</div>
			</div>
		</Modal>
	{/await}
</div>

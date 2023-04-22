<script lang="ts">
	import { hostname } from '../stores/environmentStore';
	import Fa from 'svelte-fa';
	import { faInfoCircle, faTerminal } from '@fortawesome/free-solid-svg-icons';
	import { Button, Spinner, Modal, Card, Alert } from 'flowbite-svelte';
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
		} else {
			throw new Error('Failed to fetch jobs');
		}
	};
</script>

<div class="px-4 md:px-16">
	<Button color="blue" href="/create">Create Container</Button>
	{#await fetchJobs()}
		<div class="grid h-96 place-items-center" data-testid="spinner">
			<Spinner />
		</div>
	{:then jobs}
		{#if !jobs}
			<Alert class="my-4 max-w-md">
				<span slot="icon"><Fa icon={faInfoCircle} /> </span>
				<span class="font-medium">It seems like you haven't created any containers yet...</span>
			</Alert>
		{:else}
			<div data-testid="job-list">
				{#each jobs as job}
					<Card class="my-3" href={`/container/${job.ID}`}>
						<div class="flex items-center my-1">
							<Fa icon={faTerminal} color="white" class="pr-6" />
							<span class="text-xl text-white">{job.Name}</span>
						</div>
					</Card>
				{/each}
			</div>
		{/if}
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

<script lang="ts">
	import Fa from 'svelte-fa';
	import { faInfoCircle, faTerminal } from '@fortawesome/free-solid-svg-icons';
	import { Button, Card, Alert } from 'flowbite-svelte';
	import { page } from '$app/stores';
</script>

<div class="px-4 md:px-16">
	<Button color="blue" href="/create">Create Container</Button>
	{#if !$page.data.jobs}
		<Alert class="my-4 max-w-md">
			<span slot="icon"><Fa icon={faInfoCircle} /> </span>
			<span class="font-medium">It seems like you haven't created any containers yet...</span>
		</Alert>
	{:else}
		<div data-testid="job-list">
			{#each $page.data.jobs as job}
				<Card class="my-3" href={`/container/${job.ID}`}>
					<div class="flex items-center my-1">
						<Fa icon={faTerminal} color="white" class="pr-6" />
						<span class="text-xl text-white">{job.Name}</span>
					</div>
				</Card>
			{/each}
		</div>
	{/if}
</div>

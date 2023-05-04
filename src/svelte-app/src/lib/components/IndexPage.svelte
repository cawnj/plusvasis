<script lang="ts">
	import { faInfoCircle, faTerminal } from '@fortawesome/free-solid-svg-icons';
	import { Alert, Button, Card, Modal } from 'flowbite-svelte';
	import Fa from 'svelte-fa';

	import type { PageData } from '../../routes/$types';

	export let data: PageData;
</script>

<div class="px-4 pb-4 md:px-16">
	<Button color="blue" href="/create" class="mr-2">Create Container</Button>
	{#if !data.jobs}
		<Alert class="my-4 max-w-md">
			<span slot="icon"><Fa icon={faInfoCircle} /> </span>
			<span class="font-medium">It seems like you haven't created any containers yet...</span>
		</Alert>
	{:else}
		<div data-testid="job-list">
			{#each data.jobs as job}
				<Card class="my-3" href={`/container/${job.ID}`}>
					<div class="my-1 flex items-center">
						<Fa icon={faTerminal} color="white" class="pr-6" />
						<span class="text-xl text-white">{job.Name}</span>
					</div>
				</Card>
			{/each}
		</div>
	{/if}
	{#if data.error}
		<Modal title="Error" open={true}>
			<div class="grid w-40 justify-center">
				<p class="text-base leading-relaxed text-gray-500 dark:text-gray-400">
					{data.error}
				</p>
				<div class="mt-4 flex justify-center">
					<Button color="blue" href="javascript:window.location.href=window.location.href"
						>Retry</Button
					>
				</div>
			</div>
		</Modal>
	{/if}
</div>

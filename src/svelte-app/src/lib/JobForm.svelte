<script lang="ts">
	import { JobFields } from '$lib/Types';
	import { Label, Input, Helper, Select, Toggle, Button } from 'flowbite-svelte';
	import type { Job } from '$lib/Types';
	import { currJob } from '../stores/nomadStore';

	export let type: string;

	let job: Job;
	currJob.subscribe((value) => {
		job = value;
	});
</script>

<div class="mb-6">
	{#each JobFields as { key, value }}
		{#if key !== 'containerName'}
			<div class="mb-3 mt-3">
				<Label class="block mb-2">{value.title}</Label>
				{#if value.type === 'input'}
					<Input label={key} id={key} name={key} placeholder={value.placeholder} value={job[key]} />
				{:else if value.type === 'option'}
					<Select label={key} id={key} name={key} items={value.options} value={job[key]} />
				{:else if value.type === 'toggle'}
					<Toggle label={key} id={key} name={key} checked={job[key]} />
				{/if}
				<Helper class="text-sm mt-2">{value.info}</Helper>
			</div>
		{/if}
	{/each}
</div>
<div>
	{#if type === 'update'}
		<Button color="blue" on:click={() => {}}>Update Container</Button>
	{:else if type === 'create'}
		<Button color="blue" on:click={() => {}}>Create Container</Button>
	{/if}
</div>

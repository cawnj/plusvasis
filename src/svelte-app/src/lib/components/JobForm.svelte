<script lang="ts">
	import { Button, Helper, Input, Label, Range, Select, Spinner, Toggle } from 'flowbite-svelte';
	import { onMount } from 'svelte';
	import { get } from 'svelte/store';

	import { goto } from '$app/navigation';
	import type { Job } from '$lib/types/Types';
	import { JobFields } from '$lib/types/Types';
	import { MakeJob } from '$lib/utils/MakeJob';
	import { fetchJobCreate, fetchJobUpdate } from '$lib/utils/NomadClient';

	import { currJob } from '../../stores/nomadStore';

	export let type: string;
	let loading = false;

	let job: Job;
	currJob.subscribe((value) => {
		job = value;
	});

	let cpu: number;
	let memory: number;
	onMount(() => {
		cpu = get(currJob).cpu;
		memory = get(currJob).memory;
	});

	function handleSubmit(event: Event) {
		let formData: FormData;
		if (event.target instanceof HTMLFormElement) {
			formData = new FormData(event.target);
		} else {
			console.log('Error');
			return;
		}

		const newJob: Job = MakeJob(formData);
		if (type === 'update') {
			newJob.containerName = job.containerName;
		}

		submitJob(newJob);
		reroute();
	}

	function submitJob(job: Job) {
		console.log(job);
		if (type === 'create') {
			fetchJobCreate(job);
		} else if (type === 'update') {
			fetchJobUpdate(job);
		}
	}

	function reroute() {
		loading = true;
		setTimeout(() => {
			if (type === 'create') {
				goto('/');
			} else if (type === 'update') {
				window.location.reload();
			}
		}, 1000);
	}
</script>

<form on:submit|preventDefault={(event) => handleSubmit(event)} data-testid="job-form" {type}>
	{#each JobFields as { key, value }}
		<!-- don't show containerName field on update -->
		{#if !(key === 'containerName' && type === 'update')}
			<div class="mb-3 mt-3">
				<Label class="mb-2 block">{value.title}</Label>
				{#if value.type === 'input'}
					<Input
						label={key}
						id={key}
						name={key}
						placeholder={value.placeholder}
						value={job[key]}
						autocapitalize="off"
						autocomplete="off"
						autocorrect="off"
					/>
				{:else if value.type === 'option'}
					<Select label={key} id={key} name={key} items={value.options} value={job[key]} />
				{:else if value.type === 'toggle'}
					<Toggle label={key} id={key} name={key} checked={job[key]} />
				{:else if value.type === 'cpuRange'}
					<Range
						class="mb-4 w-1/4"
						label={key}
						id={key}
						name={key}
						min="100"
						max="1000"
						step="50"
						size="lg"
						bind:value={cpu}
					/>
					<Label class="mb-2 block">Value: {cpu}</Label>
				{:else if value.type === 'memoryRange'}
					<Range
						class="mb-4 w-1/4"
						label={key}
						id={key}
						name={key}
						min="300"
						max="2000"
						step="50"
						size="lg"
						bind:value={memory}
					/>
					<Label class="mb-2 block">Value: {memory}</Label>
				{/if}
				<Helper class="mt-2 whitespace-pre-line text-sm">{value.info}</Helper>
			</div>
		{/if}
	{/each}
	<div class="pt-4">
		<Button color="blue" type="submit">
			{#if loading}
				<Spinner class="mr-3" size="4" color="white" />Loading ...
			{:else}
				{type === 'create' ? 'Create Container' : 'Update Container'}
			{/if}
		</Button>
	</div>
</form>

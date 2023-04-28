<script lang="ts">
	import { JobFields } from '$lib/Types';
	import { Label, Input, Helper, Select, Toggle, Button, Spinner, Range } from 'flowbite-svelte';
	import type { Job } from '$lib/Types';
	import { currJob } from '../stores/nomadStore';
	import { MakeJob } from '$lib/MakeJob';
	import { fetchJobCreate, fetchJobUpdate } from '$lib/NomadClient';
	import { goto } from '$app/navigation';

	export let type: string;
	let loading = false;

	let job: Job;
	currJob.subscribe((value) => {
		job = value;
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
				<Label class="block mb-2">{value.title}</Label>
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
						min="100"
						max="1000"
						step="50"
						size="lg"
						bind:value={job.cpu}
					/>
					<Label class="block mb-2">Value: {job.cpu}</Label>
				{:else if value.type === 'memoryRange'}
					<Range
						class="mb-4 w-1/4"
						label={key}
						id={key}
						min="300"
						max="2000"
						step="50"
						size="lg"
						bind:value={job.memory}
					/>
					<Label class="block mb-2">Value: {job.memory}</Label>
				{/if}
				<Helper class="text-sm mt-2 whitespace-pre-line">{value.info}</Helper>
			</div>
		{/if}
	{/each}
	<Button color="blue" type="submit">
		{#if loading}
			<Spinner class="mr-3" size="4" color="white" />Loading ...
		{:else}
			{type === 'create' ? 'Create Container' : 'Update Container'}
		{/if}
	</Button>
</form>

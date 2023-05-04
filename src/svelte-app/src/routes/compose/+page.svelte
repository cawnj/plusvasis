<script lang="ts">
	import { Button, Heading, Spinner } from 'flowbite-svelte';
	import * as yaml from 'js-yaml';

	import { goto } from '$app/navigation';
	import Editor from '$lib/components/Editor.svelte';
	import NavBar from '$lib/components/NavBar.svelte';
	import type { DockerCompose } from '$lib/types/Types';
	import { MakeJobsFromCompose } from '$lib/utils/MakeJob';
	import { fetchJobCreate } from '$lib/utils/NomadClient';

	let dockerComposeStr: string;
	let loading = false;

	function handleSubmit() {
		const dockerComposeYaml = yaml.load(dockerComposeStr);
		const dockerComposeData: DockerCompose = dockerComposeYaml as DockerCompose;
		const jobs = MakeJobsFromCompose(dockerComposeData);

		for (const job of jobs) {
			fetchJobCreate(job);
		}
		loading = true;
		setTimeout(() => {
			goto('/');
		}, 1000);
	}
</script>

<NavBar />
<div class="px-4 pb-4 md:px-16">
	<Heading tag="h3" class="mb-4 font-semibold text-white">docker-compose</Heading>
	<div class="relative">
		<Editor bind:value={dockerComposeStr} />
		<Button class="absolute bottom-4 right-4" color="blue" on:click={handleSubmit}>
			{#if loading}
				<Spinner class="mr-3" size="4" color="white" />Loading ...
			{:else}
				Submit
			{/if}
		</Button>
	</div>
</div>

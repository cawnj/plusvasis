<script lang="ts">
	import { Button, Heading } from 'flowbite-svelte';
	import * as yaml from 'js-yaml';

	import Editor from '$lib/components/Editor.svelte';
	import NavBar from '$lib/components/NavBar.svelte';
	import type { DockerCompose, Job } from '$lib/types/Types';
	import { MakeJobsFromCompose } from '$lib/utils/MakeJob';

	import { currJob } from '../../stores/nomadStore';

	// default job configuration
	currJob.set({
		shell: '/bin/sh',
		cpu: 100,
		memory: 300
	} as Job);

	let dockerComposeStr: string;

	function handleSubmit() {
		const dockerComposeYaml = yaml.load(dockerComposeStr);
		const dockerComposeData: DockerCompose = dockerComposeYaml as DockerCompose;
		const jobs = MakeJobsFromCompose(dockerComposeData);
		console.log(jobs);
	}
</script>

<NavBar />
<div class="px-4 pb-4 md:px-16">
	<Heading tag="h3" class="mb-4 font-semibold text-white">docker-compose</Heading>
	<div class="relative">
		<Editor bind:value={dockerComposeStr} />
		<Button class="absolute bottom-4 right-4" color="blue" on:click={handleSubmit}>Submit</Button>
	</div>
</div>

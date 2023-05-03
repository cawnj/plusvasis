<script lang="ts">
	import { Button, Heading } from 'flowbite-svelte';

	import Editor from '$lib/components/Editor.svelte';
	import NavBar from '$lib/components/NavBar.svelte';
	import type { Job } from '$lib/types/Types';

	import { currJob } from '../../stores/nomadStore';

	// default job configuration
	currJob.set({
		shell: '/bin/sh',
		cpu: 100,
		memory: 300
	} as Job);
</script>

<NavBar />
<div class="px-4 pb-4 md:px-16">
	<Heading tag="h2" class="mb-4 font-semibold text-white">docker-compose</Heading>
	<div class="relative">
		<Editor />
		<Button class="absolute bottom-4 right-4" color="blue">Submit</Button>
	</div>
</div>

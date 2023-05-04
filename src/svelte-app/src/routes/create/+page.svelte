<script lang="ts">
	import { Button, Hr } from 'flowbite-svelte';

	import JobForm from '$lib/components/JobForm.svelte';
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
<div class="px-4 pb-8 md:px-16">
	<JobForm type="create" />
	<div class="mb-4 py-4">
		<Hr
			width="w-64"
			divClass="inline-flex items-center w-full"
			textSpanClass="absolute left-32 px-3 font-medium text-gray-900 bg-white -translate-x-1/2 dark:text-white"
		>
			or
		</Hr>
	</div>
	<Button color="purple" href="/compose">Import from docker-compose</Button>
</div>

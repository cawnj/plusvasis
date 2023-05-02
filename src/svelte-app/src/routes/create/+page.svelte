<script lang="ts">
	import NavBar from '$lib/components/NavBar.svelte';
	import type { Job } from '$lib/types/Types';
	import JobForm from '$lib/components/JobForm.svelte';
	import { currJob } from '../../stores/nomadStore';

	// default job configuration
	currJob.set({
		shell: '/bin/sh',
		cpu: 100,
		memory: 300
	} as Job);
</script>

<NavBar />
<div class="px-4 md:px-16 pb-4">
	<JobForm type="create" />
</div>

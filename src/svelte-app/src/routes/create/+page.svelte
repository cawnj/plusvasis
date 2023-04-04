<script lang="ts">
	import NavBar from '$lib/NavBar.svelte';
	import type { Job } from '$lib/Types';
	import JobForm from '$lib/JobForm.svelte';
	import { currJob } from '../../stores/nomadStore';

	// default job configuration
	currJob.set({
		shell: '/bin/sh'
	} as Job);
</script>

<NavBar />
<div class="px-8 md:px-16 mb-4">
	<JobForm type="create" />
</div>

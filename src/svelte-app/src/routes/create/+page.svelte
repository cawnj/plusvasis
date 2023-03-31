<script lang="ts">
	import { goto } from '$app/navigation';
	import NavBar from '$lib/NavBar.svelte';
	import type { Job } from '$lib/Types';
	import { hostname } from '../../stores/environmentStore';
	import JobForm from '$lib/JobForm.svelte';
	import { currJob } from '../../stores/nomadStore';

	// default job configuration
	currJob.set({
		shell: '/bin/sh'
	} as Job);
</script>

<NavBar />
<div class="px-8 md:px-16">
	<JobForm type="create" />
</div>

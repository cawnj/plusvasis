<script lang="ts">
	import type { Job } from '$lib/Types';
	import { currJobId, currJob } from '../stores/nomadStore';
	import { hostname } from '../stores/environmentStore';
	import { Button } from 'flowbite-svelte';
	import JobForm from './JobForm.svelte';

	let jobId: string;
	let job: Job;
	currJobId.subscribe((value) => {
		jobId = value;
	});
	currJob.subscribe((value) => {
		job = value;
	});
</script>

<JobForm type="update" />

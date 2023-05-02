<script lang="ts">
	import type { Job } from '$lib/types/Types';

	import { currJob, currJobId } from '../../stores/nomadStore';
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

<script lang="ts">
	import ExecController from '$lib/ExecController.svelte';
	import { currJobId, currJob, currJobStopped } from '../stores/nomadStore';
	import { hostname } from '../stores/environmentStore';
	import type { Job } from '$lib/Types';

	let execControllerComponent: ExecController;
	let wsUrl: string;

	let jobId: string;
	let job: Job;
	let isStopped: boolean;
	currJobId.subscribe((value) => {
		jobId = value;
	});
	currJob.subscribe((value) => {
		job = value;
	});
	currJobStopped.subscribe((value) => {
		isStopped = value;
	});

	function setExecUrl() {
		const url = new URL(`${hostname}/job/${jobId}/exec`);
		url.protocol = url.protocol.replace('http', 'ws');
		url.searchParams.append('command', `["${job.shell}"]`);
		wsUrl = url.toString();
	}

	$: if (!isStopped && jobId && job) {
		setExecUrl();
	}
</script>

<ExecController bind:this={execControllerComponent} {wsUrl} />

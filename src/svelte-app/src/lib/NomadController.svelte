<script lang="ts">
	import ExecController from '$lib/ExecController.svelte';
	import { currJobId, currJob, currJobStopped } from '../stores/nomadStore';
	import { hostname } from '../stores/environmentStore';
	import type { Job } from '$lib/Types';
	import { token } from '../stores/auth';

	let execControllerComponent: ExecController;
	let wsUrl: string;

	let jobId: string;
	let job: Job;
	let isStopped: boolean;
	let authToken: string | undefined;
	currJobId.subscribe((value) => {
		jobId = value;
	});
	currJob.subscribe((value) => {
		job = value;
	});
	currJobStopped.subscribe((value) => {
		isStopped = value;
	});
	token.subscribe((value) => {
		authToken = value;
	});

	function setExecUrl() {
		const url = new URL(`${hostname}/job/${jobId}/exec`);
		url.protocol = url.protocol.replace('http', 'ws');
		url.searchParams.append('command', `["${job.shell}"]`);

		if (authToken) url.searchParams.append('access_token', authToken);

		wsUrl = url.toString();
	}

	$: if (!isStopped && jobId && job) {
		setExecUrl();
	}
</script>

<ExecController bind:this={execControllerComponent} {wsUrl} />

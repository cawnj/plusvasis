<script lang="ts">
	import ExecController from '$lib/ExecController.svelte';
	import { currJobId, currJob, alloc, task, currJobStopped } from '../stores/nomadStore';
	import type { Job } from '$lib/Types';
	import { fetchJobIdAllocations } from '$lib/NomadClient';

	let execControllerComponent: ExecController;
	let allocId: string;
	let taskName: string;

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

	// TODO: Remove references to alloc and task once LogController has a proxy too
	function getAllocExecEndpoint(json: unknown) {
		if (json?.ID && json?.TaskStates) {
			allocId = json.ID;
			taskName = Object.keys(json.TaskStates)[0];

			alloc.set(allocId);
			task.set(taskName);
		} else {
			console.log(json);
			throw new Error('Invalid JSON');
		}

		const url = new URL(`${hostname}/job/${jobId}/exec`);
		url.protocol = url.protocol.replace('http', 'ws');
		url.searchParams.append('task', taskName);
		url.searchParams.append('command', `["${job.shell}"]`);
		url.searchParams.append('tty', 'true');
		url.searchParams.append('ws_handshake', 'true');

		return url.toString();
	}

	let url: string;
	$: if (job && !isStopped && execControllerComponent) {
		fetchJobIdAllocations().then((json) => {
			url = getAllocExecEndpoint(json);
		});
	}
</script>

<ExecController bind:this={execControllerComponent} {url} />

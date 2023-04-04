<script lang="ts">
	import ExecController from '$lib/ExecController.svelte';
	import { currJobId, currJob, alloc, task } from '../stores/nomadStore';
	import { hostname } from '../stores/environmentStore';
	import { onMount } from 'svelte';
	import type { Job } from '$lib/Types';

	let execControllerComponent: ExecController;
	let allocId: string;
	let taskName: string;

	let jobId: string;
	let job: Job;
	currJobId.subscribe((value) => {
		jobId = value;
	});
	currJob.subscribe((value) => {
		job = value;
	});

	function getAllocExecEndpoint(json: unknown) {
		if (typeof json === 'object' && json !== null) {
			allocId = (json as { ID: string })['ID'];
			taskName = Object.keys((json as { TaskStates: Record<string, unknown> })['TaskStates'])[0];

			alloc.set(allocId);
			task.set(taskName);
		} else {
			throw new Error('Invalid JSON');
		}

		const url = new URL(`wss://nomad.local.cawnj.dev/v1/client/allocation/${allocId}/exec`);
		url.searchParams.append('task', taskName);
		url.searchParams.append('command', `["${job.shell}"]`);
		url.searchParams.append('tty', 'true');
		url.searchParams.append('ws_handshake', 'true');

		return url.toString();
	}

	async function fetchJobIdAllocations() {
		const url = `${hostname}/job/${jobId}/alloc`;
		const res = await fetch(url, {
			headers: {
				Authorization: `Bearer ${localStorage.getItem('token')}`
			}
		});

		if (res.ok) {
			const json = await res.json();
			const url = getAllocExecEndpoint(json);
			execControllerComponent.connectTerm(url);
		} else {
			execControllerComponent.write('Error starting container ' + jobId);
		}
	}

	onMount(async () => {
		fetchJobIdAllocations();
	});
</script>

<ExecController bind:this={execControllerComponent} />

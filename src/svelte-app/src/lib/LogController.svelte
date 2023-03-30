<script lang="ts">
	import { hostname } from '../stores/environmentStore';
	import { job, alloc, task } from '../stores/nomadStore';
	import ExecController from '$lib/ExecController.svelte';
	import { onMount } from 'svelte';

	export let jobId = '';
	export let allocId = '';
	export let taskName = '';

	let execControllerComponent: ExecController;

	job.subscribe((value) => {
		jobId = value;
	});
	alloc.subscribe((value) => {
		allocId = value;
	});
	task.subscribe((value) => {
		taskName = value;
	});

	export async function fetchLogs() {
		const url = `${hostname}/logs/${jobId}/${allocId}/${taskName}/stdout`;
		const res = await fetch(url, {
			headers: {
				Authorization: `Bearer ${localStorage.getItem('token')}`
			}
		});

		if (res.ok) {
			const json = await res.json();
			execControllerComponent.write(atob(json.data));
		} else {
			execControllerComponent.write('Error fetching container logs');
		}
	}

	onMount(async () => {
		fetchLogs();
	});
</script>

<ExecController bind:this={execControllerComponent} />

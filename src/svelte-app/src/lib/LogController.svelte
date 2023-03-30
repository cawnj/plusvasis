<script lang="ts">
	import { hostname } from '../stores/environmentStore';
	import { job, alloc, task } from '../stores/nomadStore';
	import ExecController from '$lib/ExecController.svelte';
	import { onMount } from 'svelte';
	import { error } from '@sveltejs/kit';

	export let jobId = '';
	export let allocId = '';
	export let taskName = '';
	export let stdOutLogs = '';
	export let stdErrLogs = '';

	job.subscribe((value) => {
		jobId = value;
	});
	alloc.subscribe((value) => {
		allocId = value;
	});
	task.subscribe((value) => {
		taskName = value;
	});

	export async function fetchLogs(type: string) {
		const url = `${hostname}/logs/${jobId}/${allocId}/${taskName}/${type}`;
		const res = await fetch(url, {
			headers: {
				Authorization: `Bearer ${localStorage.getItem('token')}`
			}
		});

		if (res.ok) {
			const json = await res.json();
			if (type == 'stdout') {
				stdOutLogs = atob(json.Data);
			} else {
				stdErrLogs = atob(json.Data);
			}
		} else {
			return error;
		}
	}

	onMount(async () => {
		fetchLogs('stdout');
		fetchLogs('stderr');
	});
</script>

<div class="text-white">${stdOutLogs}</div>
<div class="text-white mt-6">${stdErrLogs}</div>

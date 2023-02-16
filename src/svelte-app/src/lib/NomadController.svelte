<script lang="ts">
	import ExecController from '$lib/ExecController.svelte';

	let execControllerComponent: ExecController;
	export let getContainerClicked = false;
	export let containerRunning = false;
	export let job = '';

	export function startContainer() {
		execControllerComponent.write('Starting container...');
	}

	export function getContainers() {
		getContainerClicked = true;
	}

	export async function fetchJobId(jobId: string) {
		containerRunning = true;
		job = jobId;
		const url = 'http://localhost:8080/job/' + jobId;
		const res = await fetch(url);

		if (res.ok) {
			execControllerComponent.write('Starting container ' + jobId);
		} else {
			execControllerComponent.write('Error starting container ' + jobId);
		}
	}

	export async function fetchJobIdDelete(jobId: string) {
		const url = 'http://localhost:8080/job/' + jobId;
		const res = await fetch(url, {
			method: 'DELETE'
		});

		if (res.ok) {
			execControllerComponent.write('Stopping container ' + jobId);
		} else {
			execControllerComponent.write('Error stopping container ' + jobId);
		}
		containerRunning = false;
	}
</script>

<ExecController bind:this={execControllerComponent} />

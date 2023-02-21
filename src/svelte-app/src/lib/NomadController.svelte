<script lang="ts">
	import ExecController from '$lib/ExecController.svelte';

	let execControllerComponent: ExecController;
	export let getContainerClicked = false;
	export let containerRunning = false;
	export let job = '';

	let nomadAllocExecEndpoint = 'wss://nomad.local.cawnj.dev/v1/client/allocation/';
	let nomadAllocExecQueryParams =
		'/exec?task=server&tty=true&ws_handshake=true&command=%5B%22%2Fbin%2Fbash%22%5D';

	export function getContainers() {
		getContainerClicked = true;
	}

	export async function fetchJobId(jobId: string) {
		containerRunning = true;
		job = jobId;
		const url = `http://localhost:8080/job/${jobId}/allocations`;
		const res = await fetch(url);
		const json = await res.json();
		const allocId = json[0]['ID'];

		if (res.ok) {
			execControllerComponent.write('Starting container ' + jobId);
			execControllerComponent.connectTerm(
				nomadAllocExecEndpoint + allocId + nomadAllocExecQueryParams
			);
		} else {
			execControllerComponent.write('Error starting container ' + jobId);
		}
	}

	export async function fetchJobIdDelete(jobId: string) {
		const url = `http://localhost:8080/job/${jobId}`;
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

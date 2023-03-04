<script lang="ts">
	import { goto } from '$app/navigation';
	import ExecController from '$lib/ExecController.svelte';
	import { nomadAllocExecEndpoint, nomadAllocExecQueryParams, job } from '../stores/nomadStore';

	let execControllerComponent: ExecController;
	export let getContainerClicked = false;
	export let getContainerCreatedClicked = false;
	export let containerRunning = false;
	export let containerName = '';
	export let dockerImage = '';
	export let ports = '';

	let jsonData;

	export function getContainers() {
		getContainerClicked = true;
	}

	export function getContainerCreated() {
		getContainerCreatedClicked = true;
	}

	export async function fetchJobId(jobId: string) {
		const url = `http://localhost:8080/job/${jobId}`;
		const res = await fetch(url);

		if (res.ok) {
			execControllerComponent.write('Starting container ' + jobId);
		} else {
			execControllerComponent.write('Error starting container ' + jobId);
		}
	}

	export async function fetchJobIdAllocations(jobId: string) {
		containerRunning = true;
		job.update(() => jobId);
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

	export function createJobJson() {
		containerName = document.getElementById('containerNameInput').value;
		dockerImage = document.getElementById('dockerImageInput').value;
		ports = document.getElementById('portsInput').value;
		jsonData = {
			containerName: containerName,
			dockerImage: dockerImage,
			ports: ports,
			email: localStorage.getItem('email')
		};

		return jsonData;
	}

	export async function fetchJobCreate() {
		const url = `http://localhost:8080/jobs`;
		const json = createJobJson();
		const res = await fetch(url, {
			method: 'POST',
			body: JSON.stringify(json)
		});

		if (res.ok) {
			console.log('Container Created');
		} else {
			console.log('Error');
		}
		getContainerCreatedClicked = false;
		goto('/');
	}
</script>

<ExecController bind:this={execControllerComponent} />

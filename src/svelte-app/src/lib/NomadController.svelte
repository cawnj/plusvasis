<script lang="ts">
	import { goto } from '$app/navigation';
	import ExecController from '$lib/ExecController.svelte';
	import { job } from '../stores/nomadStore';
	import { hostname } from '../stores/environmentStore';

	let execControllerComponent: ExecController;
	export let containerName = '';
	export let dockerImage = '';

	function getAllocExecEndpoint(json: any) {
		const allocId = json[0]['ID'];
		const taskName = Object.keys(json[0]['TaskStates'])[0];
		const command = '["/bin/bash"]';

		const url = new URL(`wss://nomad.local.cawnj.dev/v1/client/allocation/${allocId}/exec`);
		url.searchParams.append('task', taskName);
		url.searchParams.append('command', command);
		url.searchParams.append('tty', 'true');
		url.searchParams.append('ws_handshake', 'true');

		return url.toString();
	}

	export async function fetchJobIdAllocations(jobId: string) {
		job.update(() => jobId);
		const url = `${hostname}/job/${jobId}/allocations`;
		const res = await fetch(url, {
			headers: {
				Authorization: `Bearer ${localStorage.getItem('token')}`
			}
		});

		if (res.ok) {
			const json = await res.json();
			const url = getAllocExecEndpoint(json);
			execControllerComponent.write('Starting container ' + jobId);
			execControllerComponent.connectTerm(url);
		} else {
			execControllerComponent.write('Error starting container ' + jobId);
		}
	}

	export async function fetchJobIdDelete(jobId: string) {
		const url = `${hostname}/job/${jobId}`;
		const res = await fetch(url, {
			method: 'DELETE',
			headers: {
				Authorization: `Bearer ${localStorage.getItem('token')}`
			}
		});

		if (res.ok) {
			execControllerComponent.write('Stopping container ' + jobId);
		} else {
			execControllerComponent.write('Error stopping container ' + jobId);
		}
	}

	export function createJobJson() {
		containerName = (<HTMLInputElement>document.getElementById('containerNameInput')).value;
		dockerImage = (<HTMLInputElement>document.getElementById('dockerImageInput')).value;
		let jsonData = {
			id: containerName + '-' + localStorage.getItem('uid'),
			containerName: containerName,
			dockerImage: dockerImage,
			user: localStorage.getItem('uid')
		};

		return jsonData;
	}

	export async function fetchJobCreate() {
		const url = `${hostname}/jobs`;
		const json = createJobJson();
		const res = await fetch(url, {
			method: 'POST',
			body: JSON.stringify(json),
			headers: {
				Authorization: `Bearer ${localStorage.getItem('token')}`
			}
		});

		if (res.ok) {
			console.log('Container Created');
		} else {
			console.log('Error');
		}
		goto('/');
	}

	export async function fetchJobUpdate(jobId: string) {
		const url = `${hostname}/job/${jobId}`;
		const json = createJobJson();
		const res = await fetch(url, {
			method: 'POST',
			body: JSON.stringify(json),
			headers: {
				Authorization: `Bearer ${localStorage.getItem('token')}`
			}
		});

		if (res.ok) {
			console.log('Container Created');
		} else {
			console.log('Error');
		}
		goto('/');
	}
</script>

<ExecController bind:this={execControllerComponent} />

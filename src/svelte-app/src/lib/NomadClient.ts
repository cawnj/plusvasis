import { hostname } from '../stores/environmentStore';
import { currJobId, currJobStopped } from '../stores/nomadStore';
import type { Job } from '$lib/Types';
import fetch from 'cross-fetch';

let jobId: string;
currJobId.subscribe((value) => {
	jobId = value;
});

export async function fetchJobCreate(job: Job) {
	const url = `${hostname}/jobs`;
	const res = await fetch(url, {
		method: 'POST',
		body: JSON.stringify(job),
		headers: {
			Authorization: `Bearer ${localStorage.getItem('token')}`
		}
	});

	if (res.ok) {
		console.log('Container Created');
	} else {
		console.log('Error');
	}
}

export async function fetchJobUpdate(job: Job) {
	const url = `${hostname}/job/${jobId}`;
	const res = await fetch(url, {
		method: 'POST',
		body: JSON.stringify(job),
		headers: {
			Authorization: `Bearer ${localStorage.getItem('token')}`
		}
	});

	if (res.ok) {
		console.log('Container Updated');
	} else {
		console.log('Error');
	}
}

export async function fetchJob(jobId: string) {
	const url = `${hostname}/job/${jobId}`;
	const res = await fetch(url, {
		headers: {
			Authorization: `Bearer ${localStorage.getItem('token')}`
		}
	});
	if (res.ok) {
		const data = await res.json();

		const job: Job = {
			user: localStorage.getItem('uid'),
			containerName: data.Name,
			dockerImage: data.TaskGroups[0].Tasks[0].Config.image,
			shell: data.Meta.shell,
			volumes: data.Meta.volumes,
			env: data.Meta.env,
			port: parseInt(data.Meta.port),
			expose: data.Meta.expose === 'true' ? true : false,
			cpu: parseInt(data.TaskGroups[0].Tasks[0].Resources.CPU),
			memory: parseInt(data.TaskGroups[0].Tasks[0].Resources.MemoryMB)
		};
		currJobStopped.set(data.Status === 'dead');
		return job;
	} else {
		console.log('Error');
		return null;
	}
}

export async function fetchJobStop() {
	const url = `${hostname}/job/${jobId}`;
	const res = await fetch(url, {
		method: 'DELETE',
		headers: {
			Authorization: `Bearer ${localStorage.getItem('token')}`
		}
	});

	if (res.ok) {
		console.log('Container Stopped');
	} else {
		console.log('Error');
	}
}

export async function fetchJobDelete() {
	const url = `${hostname}/job/${jobId}?purge=true`;
	const res = await fetch(url, {
		method: 'DELETE',
		headers: {
			Authorization: `Bearer ${localStorage.getItem('token')}`
		}
	});

	if (res.ok) {
		console.log('Container Deleted');
	} else {
		console.log('Error');
	}
}

export async function fetchJobRestart() {
	const url = `${hostname}/job/${jobId}/restart`;
	const res = await fetch(url, {
		method: 'POST',
		headers: {
			Authorization: `Bearer ${localStorage.getItem('token')}`
		}
	});

	if (res.ok) {
		console.log('Container Restarted');
	} else {
		console.log('Error');
	}
}

export async function fetchJobStart() {
	const url = `${hostname}/job/${jobId}/start`;
	const res = await fetch(url, {
		method: 'GET',
		headers: {
			Authorization: `Bearer ${localStorage.getItem('token')}`
		}
	});

	if (res.ok) {
		console.log('Container Started');
	} else {
		console.log('Error');
	}
}

export async function fetchJobIdAllocations() {
	const url = `${hostname}/job/${jobId}/alloc`;
	const res = await fetch(url, {
		headers: {
			Authorization: `Bearer ${localStorage.getItem('token')}`
		}
	});

	if (res.ok) {
		const json = await res.json();
		return json;
	} else {
		throw new Error('Failed to fetch allocations');
	}
}

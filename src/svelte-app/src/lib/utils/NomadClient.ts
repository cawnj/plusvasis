import fetch from 'cross-fetch';

import type { Job } from '$lib/types/Types';

import { token } from '../../stores/auth';
import { hostname } from '../../stores/environmentStore';
import { currJobId } from '../../stores/nomadStore';

let jobId: string;
let authToken: string | undefined;
currJobId.subscribe((value) => {
	jobId = value;
});
token.subscribe((value) => {
	authToken = value;
});

export async function fetchJobCreate(job: Job) {
	const url = `${hostname}/jobs`;
	const res = await fetch(url, {
		method: 'POST',
		body: JSON.stringify(job),
		headers: {
			Authorization: `Bearer ${authToken}`
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
			Authorization: `Bearer ${authToken}`
		}
	});

	if (res.ok) {
		console.log('Container Updated');
	} else {
		console.log('Error');
	}
}

export async function fetchJobStop() {
	const url = `${hostname}/job/${jobId}`;
	const res = await fetch(url, {
		method: 'DELETE',
		headers: {
			Authorization: `Bearer ${authToken}`
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
			Authorization: `Bearer ${authToken}`
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
			Authorization: `Bearer ${authToken}`
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
			Authorization: `Bearer ${authToken}`
		}
	});

	if (res.ok) {
		console.log('Container Started');
	} else {
		console.log('Error');
	}
}

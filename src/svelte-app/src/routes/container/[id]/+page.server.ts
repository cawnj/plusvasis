import type { PageServerLoadEvent } from './$types';
import { hostname } from '../../../stores/environmentStore';
import type { Job } from '$lib/types/Types';

export async function load({ params, cookies }: PageServerLoadEvent) {
	const token = cookies.get('token');
	const jobId = params.id;

	if (!token || !jobId) return;

	async function fetchJob() {
		const url = `${hostname}/job/${jobId}`;
		const res = await fetch(url, {
			headers: {
				Authorization: `Bearer ${token}`
			}
		});
		if (res.ok) {
			const data = await res.json();
			const job: Job = {
				// we don't care about the user here, for fetching jobs, this is handled api-side
				// we only need the uid client-side when creating a job
				user: undefined,
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
			const isStopped = data.Status === 'dead';
			return { job, isStopped };
		} else {
			console.log('Error');
			return null;
		}
	}

	try {
		const data = await fetchJob();
		const job = data?.job;
		const isStopped = data?.isStopped;
		return { job, isStopped, error: null };
	} catch (e: unknown) {
		if (e instanceof Error) {
			console.log(e);
			return { job: null, isStopped: null, error: e.message };
		}
	}
}

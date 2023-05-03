import { hostname } from '../stores/environmentStore';
import type { PageServerLoadEvent } from './$types';

export async function load({ fetch, cookies }: PageServerLoadEvent) {
	const token = cookies.get('token');
	if (!token) return;

	const fetchJobs = async () => {
		const res = await fetch(`${hostname}/jobs`, {
			headers: {
				Authorization: `Bearer ${cookies.get('token')}`
			}
		});
		if (res.ok) {
			return await res.json();
		} else {
			return [];
		}
	};
	try {
		const jobs = await fetchJobs();
		return { jobs, error: null };
	} catch (e: unknown) {
		if (e instanceof Error) {
			console.log(e);
			return { jobs: [], error: e.message };
		}
	}
}

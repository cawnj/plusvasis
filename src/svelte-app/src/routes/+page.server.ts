import type { PageServerLoadEvent } from './$types';
import { hostname } from '../stores/environmentStore';

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
	const jobs = await fetchJobs();
	return { jobs };
}

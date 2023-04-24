import type { PageServerLoadEvent } from "./$types";
import { hostname } from '../../stores/environmentStore';

export async function load({ parent }: PageServerLoadEvent) {
    const { token } = await parent();
	const fetchJobs = async () => {
		const res = await fetch(`${hostname}/jobs`, {
			headers: {
				Authorization: `Bearer ${token}`
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
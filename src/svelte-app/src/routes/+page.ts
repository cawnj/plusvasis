import { hostname } from '../stores/environmentStore';

export const load = async () => {
	console.log('Running async load func');
	console.log(`Hostname: ${hostname}`);

	async function fetchJobs() {
		const res = await fetch(`${hostname}/jobs`);
		const data = await res.json();

		if (res.ok) {
			return data;
		} else {
			throw new Error(data);
		}
	}

	return {
		jobs: fetchJobs()
	};
};

export const load = async () => {
	console.log('Running async load func');

	async function fetchData() {
		const res = await fetch('http://localhost:8080/jobs');
		const data = await res.json();

		if (res.ok) {
			return data;
		} else {
			throw new Error(data);
		}
	}

	return {
		jobs: fetchData()
	};
};

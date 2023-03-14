import { SECRET_API_KEY } from '$env/static/private';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = () => {
	function createJobJson() {
		const jsonData = {
			secret: SECRET_API_KEY
		};
		return jsonData;
	}

	async function getJWT() {
		const url = `http://localhost:8080/token`;
		const json = createJobJson();
		const res = await fetch(url, {
			method: 'POST',
			body: JSON.stringify(json)
		});

		const data = await res.json();

		if (res.ok) {
			console.log(data);
		} else {
			console.log('JWT error');
		}
	}

	return getJWT();
};

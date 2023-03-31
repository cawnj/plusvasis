<script>
	import '../app.css';
	/* eslint-disable @typescript-eslint/no-unused-vars */
	import App from '$lib/fb';
	import { onMount } from 'svelte';
	import { getAuth, onAuthStateChanged } from 'firebase/auth';
	import { goto } from '$app/navigation';

	onMount(() => {
		const auth = getAuth();
		onAuthStateChanged(auth, (user) => {
			if (user) {
				console.log('Welcome to PlusVasis');
			} else {
				goto('/login');
			}
		});
	});
</script>

<slot />

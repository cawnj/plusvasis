<script>
	import '../app.css';
	/* eslint-disable @typescript-eslint/no-unused-vars */
	import App from './fb';
	import { isLoggedIn } from '../stores/authStore';
	import { onMount } from 'svelte';
	import { getAuth, onAuthStateChanged } from 'firebase/auth';
	import { goto } from '$app/navigation';

	onMount(() => {
		const auth = getAuth();
		onAuthStateChanged(auth, (user) => {
			if (user) {
				console.log('Welcome to Continens');
				isLoggedIn.update(() => true);
			} else {
				goto('/login');
				isLoggedIn.update(() => false);
			}
		});
	});
</script>

<slot />

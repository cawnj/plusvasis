<script lang="ts">
	import { goto } from '$app/navigation';
	import { getAuth, signOut } from 'firebase/auth';
	import { onMount } from 'svelte';

	const auth = getAuth();

	onMount(() => {
		signOut(auth)
			.then(() => {
				localStorage.removeItem('token');
				goto('/login');
			})
			.catch((error) => {
				console.error(error);
			});
	});
</script>

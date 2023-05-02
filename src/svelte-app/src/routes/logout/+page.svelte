<script lang="ts">
	import { goto } from '$app/navigation';
	import { signOut } from '../../stores/auth';
	import { onMount } from 'svelte';

	onMount(() => {
		signOut()
			.then(() => {
				goto('/login');
			})
			.catch((error) => {
				console.error(error);
			});
	});
</script>

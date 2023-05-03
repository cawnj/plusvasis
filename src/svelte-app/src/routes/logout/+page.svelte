<script lang="ts">
	import { onMount } from 'svelte';

	import { goto } from '$app/navigation';

	import { signOut } from '../../stores/auth';

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

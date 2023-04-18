<script lang="ts">
	import { goto } from '$app/navigation';
	import { getAuth, signOut } from 'firebase/auth';

	const auth = getAuth();
	signOut(auth)
		.then(() => {
			goto('/login');
		})
		.catch((error) => {
			console.error(error);
		});
</script>

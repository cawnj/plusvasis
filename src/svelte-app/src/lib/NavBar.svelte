<script>
	import logo from '$lib/assets/logo.png';
	import { goto } from '$app/navigation';
	import { getAuth, signOut } from 'firebase/auth';
	import { isLoggedIn } from '../stores/authStore';

	const auth = getAuth();
	function logout() {
		signOut(auth)
			.then(() => {
				localStorage.removeItem('uid');
				goto('/login');
			})
			.catch((error) => {
				console.error(error);
			});
	}
</script>

<div>
	<div>
		<nav class="container px-6 py-8 mx-auto md:flex md:justify-between md:items-center">
			<a href="/">
				<div class="flex items-center justify-between">
					<img alt="The project logo" src={logo} class="mr-3 h-6 sm:h-9 float-left" />
					<span class="text-xl font-bold text-white md:text-2xl hover:text-blue-400"
						>Continens
					</span>
				</div>
			</a>

			<div
				class="flex-col mt-8 space-y-4 md:flex md:space-y-0 md:flex-row md:items-center md:space-x-10 md:mt-0"
			>
				<a class="text-white hover:text-blue-400" href="/">Home</a>
				<a class="text-white hover:text-blue-400" href="/">About</a>
				{#if $isLoggedIn}
					<a
						class="text-white hover:text-blue-400"
						on:click|preventDefault={logout}
						target="_blank"
						href="/">Sign Out</a
					>
				{/if}
			</div>
		</nav>
	</div>
</div>

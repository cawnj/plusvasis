<script lang="ts">
	import { goto } from '$app/navigation';
	import { Card, Button, Label, Input } from 'flowbite-svelte';
	import {
		getAuth,
		signInWithEmailAndPassword,
		createUserWithEmailAndPassword
	} from 'firebase/auth';
	import logo from '$lib/assets/logo.png';
	import app from '$lib/fb';

	export let title: string;
	let email: string;
	let password: string;

	const auth = getAuth(app);

	function login(event: Event) {
		if (event.target instanceof HTMLFormElement) {
			const formData = new FormData(event.target);
			email = formData.get('email') as string;
			password = formData.get('password') as string;
		} else {
			return;
		}

		if (title === 'Login') {
			signInWithEmailAndPassword(auth, email, password)
				.then((userCredential) => {
					const user = userCredential.user;
					localStorage.setItem('uid', user.uid);
					user.getIdToken().then((token) => {
						localStorage.setItem('token', token);
					});
					goto('/');
				})
				.catch((error) => {
					const errorCode = error.code;
					const errorMessage = error.message;
					console.log('error code:', errorCode, ' error msg: ', errorMessage);
				});
		} else {
			createUserWithEmailAndPassword(auth, email, password)
				.then((userCredential) => {
					const user = userCredential.user;
					console.log(user);
					goto('/');
				})
				.catch((error) => {
					const errorCode = error.code;
					const errorMessage = error.message;
					console.log('error code:', errorCode, ' error msg: ', errorMessage);
				});
		}
	}
</script>

<Card>
	<form class="flex flex-col space-y-6" on:submit|preventDefault={login}>
		<div class="flex items-center">
			<img alt="The project logo" src={logo} class="mr-3 h-6 sm:h-9 float-left" />
			<h5 class="text-xl font-medium text-gray-900 dark:text-white">PlusVasis {title}</h5>
		</div>
		<Label class="space-y-2">
			<span>Email</span>
			<Input type="email" name="email" placeholder="name@company.com" required />
		</Label>
		<Label class="space-y-2">
			<span>Your password</span>
			<Input type="password" name="password" placeholder="•••••" required />
		</Label>
		{#if title === 'Login'}
			<Button type="submit" class="w-full">Login to your account</Button>
		{:else}
			<Button type="submit" class="w-full">Create your account</Button>
		{/if}
		{#if title === 'Login'}
			<div class="text-sm font-medium text-gray-500 dark:text-gray-300">
				Not registered? <a href="/signup" class="text-blue-700 hover:underline dark:text-blue-500"
					>Create account</a
				>
			</div>
		{/if}
		<div class="text-sm font-medium text-gray-500 dark:text-gray-300">
			We'll never share your details with anyone else.
		</div>
	</form>
</Card>

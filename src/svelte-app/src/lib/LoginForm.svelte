<script lang="ts">
	import { goto } from '$app/navigation';
	import { Card, Button, Label, Input, Alert } from 'flowbite-svelte';
	import {
		getAuth,
		signInWithEmailAndPassword,
		createUserWithEmailAndPassword
	} from 'firebase/auth';
	import logo from '$lib/assets/logo.png';
	import app from '$lib/fb';
	import { faArrowLeft, faExclamationTriangle } from '@fortawesome/free-solid-svg-icons';
	import Fa from 'svelte-fa';

	export let title: string;
	let email: string;
	let password: string;
	let errorCode: string | null = null;

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
					errorCode = error.code;
				});
		} else {
			createUserWithEmailAndPassword(auth, email, password)
				.then((userCredential) => {
					const user = userCredential.user;
					console.log(user);
					goto('/');
				})
				.catch((error) => {
					errorCode = error.code;
				});
		}
	}
</script>

<Card>
	<form class="flex flex-col space-y-6" on:submit|preventDefault={login}>
		<div class="flex items-center">
			{#if title === 'Login'}
				<img alt="The project logo" src={logo} class="mr-3 h-8 w-8" />
			{:else}
				<Button pill={true} class="!p-2 mr-3 h-8 w-8" on:click={() => history.back()}>
					<Fa icon={faArrowLeft} size="lg" />
				</Button>
			{/if}
			<h5 class="text-xl font-medium text-gray-900 dark:text-white">PlusVasis {title}</h5>
		</div>
		<Label class="space-y-2">
			<span>Email</span>
			<Input type="email" name="email" placeholder="name@example.com" required />
		</Label>
		<Label class="space-y-2">
			<span>Your password</span>
			<Input type="password" name="password" placeholder="••••••••••••" required />
		</Label>
		{#if errorCode}
			<Alert color="none" class="bg-red-100 text-red-600 border-red-800 !py-3">
				<span slot="icon">
					<Fa icon={faExclamationTriangle} class="mr-2" />
				</span>
				<span>
					{errorCode}
				</span>
			</Alert>
		{/if}

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

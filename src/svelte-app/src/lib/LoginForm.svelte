<script lang="ts">
	import { Card, Button, Label, Input, Alert, Hr, A } from 'flowbite-svelte';
	import { signInWithGithub, signInWithEmail, createUserWithEmail } from '../stores/auth';
	import logo from '$lib/assets/logo.png';
	import { faArrowLeft, faExclamationTriangle } from '@fortawesome/free-solid-svg-icons';
	import Fa from 'svelte-fa';
	import { LoginButton } from 'svelte-auth-ui';
	import { goto } from '$app/navigation';

	export let title: string;
	let errorCode: string | null = null;
	let loading = false;

	async function handleFormSubmit(event: Event) {
		let email: string;
		let password: string;
		if (event.target instanceof HTMLFormElement) {
			const formData = new FormData(event.target);
			email = formData.get('email') as string;
			password = formData.get('password') as string;
		} else {
			errorCode = 'auth/invalid-form-input';
			return;
		}

		try {
			if (title === 'Login') {
				await signInWithEmail(email, password);
			} else {
				await createUserWithEmail(email, password);
			}
			goto('/');
		} catch (error: any) {
			errorCode = error.code;
		}
	}

	async function handleGithubLogin() {
		loading = true;
		try {
			await signInWithGithub();
			goto('/');
		} catch (error: any) {
			errorCode = error.code;
		} finally {
			loading = false;
		}
	}
</script>

<Card>
	<form class="flex flex-col space-y-6" on:submit|preventDefault={handleFormSubmit}>
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
		<div class="flex flex-col space-y-6 py-2">
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
				<Button type="submit" class="w-full"><span class="text-base">Login with email</span></Button
				>
				<Hr class="my-2" width="w-64">or</Hr>
				<div class="space-y-4">
					<LoginButton
						provider="github"
						{loading}
						withLoader
						on:click={handleGithubLogin}
						class="w-full h-full"
					/>
				</div>
			{:else}
				<Button type="submit" class="w-full"
					><span class="text-base">Sign up with email</span></Button
				>
			{/if}
		</div>
		<div class="flex flex-col space-y-2">
			{#if title === 'Login'}
				<div class="text-sm font-medium text-gray-500 dark:text-gray-300">
					Not registered? <A href="/signup" class="text-blue-700 hover:underline dark:text-blue-500"
						>Create an account</A
					>
				</div>
			{/if}
			<div class="text-sm font-medium text-gray-500 dark:text-gray-300">
				We'll never share your details with anyone else.
			</div>
		</div>
	</form>
</Card>

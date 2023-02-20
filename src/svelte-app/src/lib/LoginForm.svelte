<script>
	import { goto } from '$app/navigation';
	import logo from '$lib/assets/logo.png';
	import {
		getAuth,
		signInWithEmailAndPassword,
		createUserWithEmailAndPassword
	} from 'firebase/auth';
	export let title;
	const auth = getAuth();
	function login() {
		let email = document.getElementById('emailInput').value;
		let password = document.getElementById('passInput').value;
		if (title == 'Login') {
			signInWithEmailAndPassword(auth, email, password)
				.then((userCredential) => {
					// Signed in
					const user = userCredential.user;
					localStorage.setItem('uid', user.uid);
					localStorage.setItem('isLoggedIn', true);
					goto('/');
				})
				.catch((error) => {
					const errorCode = error.code;
					const errorMessage = error.message;
					console.log('error code:', +errorCode + ' error msg: ' + errorMessage);
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
					console.log('error code:', +errorCode + ' error msg: ' + errorMessage);
				});
		}
	}
</script>

<div class="login">
	<div
		class="w-full max-w-sm p-4 bg-white border border-gray-200 rounded-lg shadow sm:p-6 md:p-8 dark:bg-gray-800 dark:border-gray-700"
	>
		<div class="card-body login-form">
			<img alt="The project logo" src={logo} class="mr-3 h-6 sm:h-9 float-left" />
			<h5 class="text-xl font-medium text-gray-900 dark:text-white">Continens {title}</h5>
			<form on:submit|preventDefault={login}>
				<div class="mb-3 mt-3">
					<label
						for="emailInput"
						class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
						>Email address</label
					>
					<input
						type="email"
						class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
						id="emailInput"
						aria-describedby="emailHelp"
						placeholder="Email Address"
					/>
				</div>
				<div class="mb-3">
					<label
						for="passInput"
						class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Password</label
					>
					<input
						type="password"
						class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
						id="passInput"
						placeholder="Password"
					/>
				</div>
				<button
					type="submit"
					class="mb-4 mt-4 w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
					>Login to your account</button
				>
			</form>
			{#if title == 'Login'}
				<div class="text-sm font-medium text-gray-500 dark:text-gray-300">
					Not registered? <a href="/signup" class="text-blue-700 hover:underline dark:text-blue-500"
						>Create account</a
					>
				</div>
			{/if}
			<div id="emailHelp" class="mt-2 text-sm font-medium text-gray-500 dark:text-gray-300">
				We'll never share your details with anyone else.
			</div>
			{#if title != 'Login'}
				<div class="text-sm font-medium text-gray-500 dark:text-gray-300">
					<a href="/login" class="mt-2 text-blue-700 hover:underline dark:text-blue-500">Back</a>
				</div>
			{/if}
		</div>
	</div>
</div>

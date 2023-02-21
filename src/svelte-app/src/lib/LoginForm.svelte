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

<div>
	<div class="div-login">
		<div class="card-body login-form">
			<div class="flex items-center">
				<img alt="The project logo" src={logo} class="mr-3 h-6 sm:h-9 float-left" />
				<h5 class="h5">Continens {title}</h5>
			</div>
			<form on:submit|preventDefault={login}>
				<div class="mb-3 mt-3">
					<label for="emailInput" class="txt-input-label">Email address</label>
					<input
						type="email"
						class="txt-input"
						id="emailInput"
						aria-describedby="emailHelp"
						placeholder="Email Address"
					/>
				</div>
				<div class="mb-3">
					<label for="passInput" class="txt-input-label">Password</label>
					<input type="password" class="txt-input" id="passInput" placeholder="Password" />
				</div>
				<button type="submit" class="btn-blue-submit">Login to your account</button>
			</form>
			{#if title == 'Login'}
				<div class="div-txt">
					Not registered? <a href="/signup" class="href-login">Create account</a>
				</div>
			{/if}
			<div id="emailHelp" class="mt-2 div-txt">
				We'll never share your details with anyone else.
			</div>
			{#if title != 'Login'}
				<div class="div-txt">
					<a href="/login" class="mt-2 href-login">Back</a>
				</div>
			{/if}
		</div>
	</div>
</div>

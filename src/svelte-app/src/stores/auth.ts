import cookie from 'cookie';
import { browser } from '$app/environment';
import {
	GithubAuthProvider,
	type User,
	signInWithPopup,
	signInWithEmailAndPassword,
	createUserWithEmailAndPassword
} from 'firebase/auth';
import { writable } from 'svelte/store';
import { auth } from '$lib/firebase/client';

export const user = writable<User | null>(null);
export const token = writable<string | null>(null);

export async function signOut() {
	return auth.signOut();
}

export async function signInWithGithub() {
	await signInWithPopup(auth, new GithubAuthProvider());
}

export async function signInWithEmail(email: string, password: string) {
	await signInWithEmailAndPassword(auth, email, password);
}

export async function createUserWithEmail(email: string, password: string) {
	await createUserWithEmailAndPassword(auth, email, password);
}

if (browser) {
	auth.onIdTokenChanged(async (newUser) => {
		const token = newUser ? await newUser?.getIdToken() : undefined;
		document.cookie = cookie.serialize('token', token ?? '', {
			path: '/',
			maxAge: token ? undefined : 0
		});
		user.set(newUser);
	});

	// refresh the ID token every 10 minutes
	setInterval(async () => {
		if (auth.currentUser) {
			await auth.currentUser.getIdToken(true);
		}
	}, 10 * 60 * 1000);
}

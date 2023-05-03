import { redirect } from '@sveltejs/kit';

import { auth } from '$lib/firebase/admin';

import type { LayoutServerLoadEvent } from './$types';

export async function load({ cookies, url }: LayoutServerLoadEvent) {
	const isAuthRoute = url.pathname == '/login' || url.pathname == '/signup';
	try {
		const token = cookies.get('token');
		if (!token && !isAuthRoute) {
			throw redirect(307, '/login');
		}
		const user = token ? await auth.verifyIdToken(token) : null;
		return {
			uid: user?.uid,
			token: token
		};
	} catch {
		// The token is set but invalid or expired
		cookies.set('token', '', { maxAge: -1 });
		throw redirect(307, '/login');
	}
}

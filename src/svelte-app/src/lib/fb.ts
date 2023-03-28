// Import the functions you need from the SDKs you need
import { initializeApp } from 'firebase/app';
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
const firebaseConfig = {
	apiKey: 'AIzaSyDmVW0FZOfgAlKsui0mYMT2g2WZBUVQLks',
	authDomain: 'plusvasis-auth.firebaseapp.com',
	projectId: 'plusvasis-auth',
	storageBucket: 'plusvasis-auth.appspot.com',
	messagingSenderId: '698224470064',
	appId: '1:698224470064:web:72faaa97b7a670ecbc25f5'
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);

export default app;

import { goto } from '$app/navigation';
import { getAuth, signOut } from 'firebase/auth';

export function logout() {
	const auth = getAuth();
	signOut(auth)
		.then(() => {
			localStorage.removeItem('uid');
			goto('/login');
		})
		.catch((error) => {
			console.error(error);
		});
}

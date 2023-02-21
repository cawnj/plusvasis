// Import the functions you need from the SDKs you need
import { initializeApp } from 'firebase/app';
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
const firebaseConfig = {
	apiKey: 'AIzaSyBFdSMSsCOpX9BDYlhxotMihiTmjsJwem8',
	authDomain: 'continens-auth.firebaseapp.com',
	projectId: 'continens-auth',
	storageBucket: 'continens-auth.appspot.com',
	messagingSenderId: '289486958559',
	appId: '1:289486958559:web:7f6f480407893357f99e6a'
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);

export default app;

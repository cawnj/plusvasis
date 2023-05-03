import { fireEvent, render, screen } from '@testing-library/svelte';
// Mock firebase/auth
import { createUserWithEmailAndPassword, getAuth, signInWithEmailAndPassword } from 'firebase/auth';
import { describe, expect, it, vi } from 'vitest';

import LoginForm from './LoginForm.svelte';
vi.mock('firebase/auth', async () => {
	const actual = await vi.importActual('firebase/auth');
	if (!actual) {
		throw new Error('Could not import actual firebase/auth');
	}

	const getAuth = vi.fn(() => ({}));
	const signInWithEmailAndPassword = vi.fn(() =>
		Promise.resolve({
			user: {
				uid: 'test',
				getIdToken: () => Promise.resolve('test')
			}
		})
	);
	const createUserWithEmailAndPassword = vi.fn();

	return {
		...actual,
		getAuth,
		signInWithEmailAndPassword,
		createUserWithEmailAndPassword
	};
});

describe('rendering', () => {
	it('should render the correct title', () => {
		render(LoginForm, { title: 'Login' });
		const loginTitle = screen.getByText('PlusVasis Login');
		expect(loginTitle).toBeInTheDocument();

		render(LoginForm, { title: 'Register' });
		const registerTitle = screen.getByText('PlusVasis Register');
		expect(registerTitle).toBeInTheDocument();
	});
});

describe('form submission', () => {
	it('should call the login function with form data on form submission', () => {
		render(LoginForm, { title: 'Login' });

		const emailInput = screen.getByPlaceholderText('name@example.com');
		const passwordInput = screen.getByPlaceholderText('••••••••••••');
		const submitButton = screen.getByText('Login with email');

		fireEvent.input(emailInput, { target: { value: 'test@example.com' } });
		fireEvent.input(passwordInput, { target: { value: 'password' } });
		fireEvent.click(submitButton);

		expect(getAuth).toBeCalled();
		expect(signInWithEmailAndPassword).toBeCalledWith(
			expect.anything(),
			'test@example.com',
			'password'
		);
		expect(createUserWithEmailAndPassword).not.toBeCalled();
	});
});

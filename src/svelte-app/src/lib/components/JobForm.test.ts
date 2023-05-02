import type { User } from '@firebase/auth';
import { fireEvent, render } from '@testing-library/svelte';
import { describe, expect, it } from 'vitest';

import { server } from '../../mocks/setup';
import { user } from '../../stores/auth';
import JobForm from './JobForm.svelte';
beforeAll(() => server.listen({ onUnhandledRequest: 'error' }));
afterEach(() => server.resetHandlers());
afterAll(() => server.close());

describe('JobForm', () => {
	it('renders the form', () => {
		const { container } = render(JobForm, {
			props: { type: 'create' }
		});
		expect(container.querySelector('form')).not.toBeNull();
	});

	it('submits job on form submit', async () => {
		const component = render(JobForm, {
			props: { type: 'create' }
		});

		// mock firebase user, uid is required for job submission
		const mockUser: User = {
			uid: 'test'
		};
		user.set(mockUser);

		const formData = new FormData();
		formData.append('containerName', 'test-container');

		await fireEvent.submit(component.container.querySelector('form'), {
			target: { formData }
		});

		expect(component.getByRole('button')).toHaveTextContent('Loading ...');
		expect(component.getByRole('status')).toBeInTheDocument();
	});
});

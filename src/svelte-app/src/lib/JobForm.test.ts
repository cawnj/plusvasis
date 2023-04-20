import { describe, it, expect } from 'vitest';
import { render, fireEvent } from '@testing-library/svelte';
import JobForm from './JobForm.svelte';

import { server } from '../mocks/setup';
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

		const formData = new FormData();
		formData.append('containerName', 'test-container');

		await fireEvent.submit(component.container.querySelector('form'), {
			target: { formData }
		});

		expect(component.getByRole('button')).toHaveTextContent('Loading ...');
		expect(component.getByRole('status')).toBeInTheDocument();
	});
});

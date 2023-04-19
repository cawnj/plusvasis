import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import { currJobId, currJob } from '../stores/nomadStore';
import LogController from './LogController.svelte';
import type { Job } from '$lib/Types';

describe('LogController', () => {
	it('should render component', async () => {
		// Initialize the Nomad store values used by the component.
		currJobId.set('job123');
		currJob.set({
			containerName: 'job123'
		} as Job);

		const { getByText } = render(LogController);

		// Verify that the component is rendered with the initial state.
		expect(getByText('stdout')).toBeTruthy();

		// Wait for the logs to be fetched and displayed.
		await new Promise((resolve) => setTimeout(resolve, 1000));

		expect(getByText('stdout')).toBeTruthy();
	});
});

import { render } from '@testing-library/svelte';
import { describe, expect, it, vi } from 'vitest';

import type { Job } from '$lib/types/Types';

import { currJob, currJobId } from '../../stores/nomadStore';
import LogController from './LogController.svelte';

vi.mock('$lib/utils/StreamLogs', () => {
	const actual = vi.importActual('$lib/utils/StreamLogs');
	return {
		...actual,
		getStream: vi.fn(() => Promise.resolve(new ReadableStream()))
	};
});

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

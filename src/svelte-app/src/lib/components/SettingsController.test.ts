import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import SettingsController from './SettingsController.svelte';

describe('SettingsController', () => {
	it('should render JobForm', async () => {
		const { getByTestId } = render(SettingsController, {
			props: {}
		});

		const jobForm = getByTestId('job-form');
		expect(jobForm).toBeInTheDocument();
		expect(jobForm).toHaveAttribute('type', 'update');
	});
});

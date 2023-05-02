import { render } from '@testing-library/svelte';
import { describe, expect, it } from 'vitest';

import type { Job } from '$lib/types/Types';

import { currJob } from '../../stores/nomadStore';
import ContainerPage from './ContainerPage.svelte';

const mockJob: Job = {
	user: 'test',
	containerName: 'test',
	dockerImage: 'test',
	shell: 'test',
	volumes: [],
	env: [],
	port: 0,
	expose: false,
	cpu: 100,
	memory: 300
};
currJob.set(mockJob);

describe('ContainerPage', () => {
	it('renders without errors', () => {
		const { container } = render(ContainerPage);
		expect(container).toBeDefined();
	});

	it('does not display the link to the container if job.expose is false', async () => {
		const { queryByText } = render(ContainerPage);
		await new Promise((r) => setTimeout(r, 1000)); // wait for fetchAndSetJob to resolve
		expect(queryByText('https://123.plusvasis.xyz')).not.toBeInTheDocument();
	});

	it('displays error modal when fetching job data fails', async () => {
		const { getByText } = render(ContainerPage);
		await new Promise((resolve) => setTimeout(resolve, 1000)); // simulate delay in fetching data
		expect(getByText('Error')).toBeInTheDocument();
	});
});

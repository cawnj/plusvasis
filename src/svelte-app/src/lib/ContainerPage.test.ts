import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import ContainerPage from './ContainerPage.svelte';

describe('ContainerPage', () => {
	it('renders without errors', () => {
		const { container } = render(ContainerPage);
		expect(container).toBeDefined();
	});

	it('displays loading spinner while fetching job data', async () => {
		const { getByTestId } = render(ContainerPage);
		expect(getByTestId('loading-spinner')).toBeInTheDocument();
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

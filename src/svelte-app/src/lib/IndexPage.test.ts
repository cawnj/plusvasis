import { describe, it, expect, vi } from 'vitest';
import { render, cleanup, waitFor, fireEvent } from '@testing-library/svelte';
import IndexPage from './IndexPage.svelte';
import { tick } from 'svelte';

describe('IndexPage', () => {
	afterEach(cleanup);

	it('should render a "Create Container" button', () => {
		const { getByText } = render(IndexPage);
		const button = getByText('Create Container');
		expect(button).toBeInTheDocument();
	});

	it('should navigate to the Create page when the "Create Container" button is clicked', async () => {
		// Render the component
		const { getByText } = render(IndexPage);

		// Click the "Create Container" button
		fireEvent.click(getByText('Create Container'));

		// Wait for the next page to render
		await tick();
	});

	it('should show a message if there are no jobs', async () => {
		// Mock the fetch function to return an empty list of jobs
		global.fetch = vi.fn(() =>
			Promise.resolve({
				ok: true,
				json: () => Promise.resolve()
			})
		);

		const { getByText } = render(IndexPage);

		// Wait for the message to appear
		await waitFor(() => {
			expect(
				getByText("It seems like you haven't created any containers yet...")
			).toBeInTheDocument();
		});
	});

	it('should show an error modal if the server returns an error', async () => {
		// Mock the fetch function to return an error
		global.fetch = vi.fn(() => Promise.reject(new Error('Failed to fetch jobs')));

		const { getByText } = render(IndexPage);

		// Wait for the error modal to appear
		await waitFor(() => {
			expect(getByText('Error')).toBeInTheDocument();
			expect(getByText('Failed to fetch jobs')).toBeInTheDocument();
		});
	});

	it('should show a list of jobs if there are jobs', async () => {
		// Mock the fetch function to return a list of jobs
		global.fetch = vi.fn(() =>
			Promise.resolve({
				ok: true,
				json: () =>
					Promise.resolve([
						{ ID: 1, Name: 'Job 1' },
						{ ID: 2, Name: 'Job 2' }
					])
			})
		);

		const { getByText } = render(IndexPage);

		// Wait for the jobs to appear
		await waitFor(() => {
			expect(getByText('Job 1')).toBeInTheDocument();
			expect(getByText('Job 2')).toBeInTheDocument();
		});
	});

	it('should show a spinner while fetching jobs', async () => {
		// Mock the fetch function to return a list of jobs after a delay
		global.fetch = vi.fn(
			() =>
				new Promise((resolve) =>
					setTimeout(
						() =>
							resolve({
								ok: true,
								json: () =>
									Promise.resolve([
										{ ID: 1, Name: 'Job 1' },
										{ ID: 2, Name: 'Job 2' }
									])
							}),
						500
					)
				)
		);

		const { getByTestId } = render(IndexPage);

		// Wait for the spinner to appear
		await waitFor(() => {
			expect(getByTestId('spinner')).toBeInTheDocument();
		});

		// Wait for the jobs to appear
		await waitFor(() => {
			expect(getByTestId('job-list')).toBeInTheDocument();
		});
	});
});

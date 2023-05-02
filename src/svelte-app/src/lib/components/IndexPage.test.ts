import { describe, it, expect } from 'vitest';
import { render, cleanup, waitFor, fireEvent } from '@testing-library/svelte';
import IndexPage from './IndexPage.svelte';
import { tick } from 'svelte';
import type { PageData } from '../routes/$types';

const defaultData: PageData = {
	uid: 'test',
	token: 'test',
	jobs: null,
	error: null
};

describe('IndexPage', () => {
	afterEach(cleanup);

	it('should render a "Create Container" button', () => {
		const { getByText } = render(IndexPage, { props: { data: defaultData } });
		const button = getByText('Create Container');
		expect(button).toBeInTheDocument();
	});

	it('should navigate to the Create page when the "Create Container" button is clicked', async () => {
		// Render the component
		const { getByText } = render(IndexPage, { props: { data: defaultData } });

		// Click the "Create Container" button
		fireEvent.click(getByText('Create Container'));

		// Wait for the next page to render
		await tick();
	});

	it('should show a message if there are no jobs', async () => {
		const { getByText } = render(IndexPage, { props: { data: defaultData } });

		// Wait for the message to appear
		await waitFor(() => {
			expect(
				getByText("It seems like you haven't created any containers yet...")
			).toBeInTheDocument();
		});
	});

	it('should show an error modal if the server returns an error', async () => {
		// Mock ssrFetch data
		const data: PageData = {
			uid: 'test',
			token: 'test',
			jobs: [],
			error: 'some failure'
		};

		const { getByText } = render(IndexPage, { props: { data: data } });

		// Wait for the error modal to appear
		await waitFor(() => {
			expect(getByText('Error')).toBeInTheDocument();
		});
	});

	it('should show a list of jobs if there are jobs', async () => {
		// Mock ssrFetch data
		const data: PageData = {
			uid: 'test',
			token: 'test',
			jobs: [
				{
					ID: 'test',
					Name: 'test'
				}
			],
			error: null
		};

		const { getByText } = render(IndexPage, { props: { data: data } });

		// Wait for the jobs to appear
		await waitFor(() => {
			expect(getByText('test')).toBeInTheDocument();
		});
	});
});

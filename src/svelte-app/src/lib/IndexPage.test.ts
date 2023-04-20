import { describe, it, expect } from 'vitest';
import { render, cleanup } from '@testing-library/svelte';
import IndexPage from './IndexPage.svelte';

describe('IndexPage', () => {
	afterEach(cleanup);

	it('should render a "Create Container" button', () => {
		const { getByText } = render(IndexPage);
		const button = getByText('Create Container');
		expect(button).toBeInTheDocument();
	});
});

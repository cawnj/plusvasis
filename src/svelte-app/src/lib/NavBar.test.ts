import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import NavBar from './NavBar.svelte';

describe('NavBar', () => {
	it('should render', () => {
		const { getByAltText, getByText, getByTestId } = render(NavBar);

		expect(getByAltText('PlusVasis Logo')).toBeInTheDocument();
		expect(getByText('PlusVasis')).toBeInTheDocument();
		expect(getByTestId('navbar-ul')).not.toBeVisible();
	});
});

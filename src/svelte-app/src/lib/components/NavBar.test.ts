import { render } from '@testing-library/svelte';
import { describe, expect, it } from 'vitest';

import NavBar from './NavBar.svelte';

describe('NavBar', () => {
	it('should render', () => {
		const { getByAltText, getByText, getByTestId } = render(NavBar);

		expect(getByAltText('PlusVasis Logo')).toBeInTheDocument();
		expect(getByText('PlusVasis')).toBeInTheDocument();
		expect(getByTestId('navbar-ul')).not.toBeVisible();
	});
});

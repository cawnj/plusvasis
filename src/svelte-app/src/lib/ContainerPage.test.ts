import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import ContainerPage from './ContainerPage.svelte';

describe('ContainerPage', () => {
	it('renders without errors', () => {
		const { container } = render(ContainerPage);
		expect(container).toBeDefined();
	});
});

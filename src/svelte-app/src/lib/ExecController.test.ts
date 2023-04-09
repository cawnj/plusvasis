import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import ExecController from './ExecController.svelte';

// Mock the matchMedia function
window.matchMedia = () => ({
	addListener: (listener: any) => {
		// You can add some code here to handle the listener
	},
	removeListener: (listener: any) => {
		// You can add some code here to handle the listener
	}
});

describe('ExecController', () => {
	it('should render', () => {
		const { container } = render(ExecController);
		expect(container.querySelector('#terminal')).not.toBeNull();
	});
});

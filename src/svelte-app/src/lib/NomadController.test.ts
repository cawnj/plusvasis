import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import NomadController from './NomadController.svelte';
import { alloc, currJobId } from '../stores/nomadStore';

// Mock the matchMedia function
window.matchMedia = () => ({
	addListener: (listener: any) => {
		// You can add some code here to handle the listener
	},
	removeListener: (listener: any) => {
		// You can add some code here to handle the listener
	}
});

describe('NomadController', () => {
	it('should render', () => {
		alloc.set('alloc-id-123');
		currJobId.set('job123');
		const { getByTestId } = render(NomadController);
		expect(getByTestId('exec-controller')).toBeInTheDocument();
	});
});

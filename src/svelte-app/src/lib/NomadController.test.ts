import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import NomadController from './NomadController.svelte';
import { alloc, currJobId } from '../stores/nomadStore';

// need to add the following to any tests that import a component
// that fetches during onMount, like NomadController
// probably a better way to do this, but this works for now
// to confirm this, run api locally and check if any requests are made to it
import { server } from '../mocks/setup';
beforeAll(() => server.listen({ onUnhandledRequest: 'error' }));
afterEach(() => server.resetHandlers());
afterAll(() => server.close());

describe('NomadController', () => {
	it('should render', () => {
		alloc.set('alloc-id-123');
		currJobId.set('job123');
		const { getByTestId } = render(NomadController);
		expect(getByTestId('exec-controller')).toBeInTheDocument();
	});
});

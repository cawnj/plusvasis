import { render } from '@testing-library/svelte';
import { describe, expect, it } from 'vitest';

import ExecController from './ExecController.svelte';

describe('ExecController', () => {
	it('should render', () => {
		const { container } = render(ExecController, {
			props: {
				wsUrl: ''
			}
		});

		expect(container.querySelector('#terminal')).not.toBeNull();
	});
});

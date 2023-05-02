import { fa0 } from '@fortawesome/free-solid-svg-icons';
import { render, screen } from '@testing-library/svelte';
import type { ComponentProps } from 'svelte';
import html from 'svelte-htm';
import { describe, expect, it } from 'vitest';

import Tabs from './Tabs.svelte';

describe('prop test', () => {
	it("doesn't pass prop", () => {
		render(Tabs);
		expect(screen.getByText('No tabs')).toBeInTheDocument();
	});

	it('passes prop with one tab', () => {
		const tabs: ComponentProps<Tabs> = {
			tabs: [
				{
					name: 'Tab1',
					component: html``,
					icon: fa0
				}
			]
		};
		render(Tabs, tabs);
		expect(screen.getByText('Tab1')).toBeInTheDocument();
	});
});

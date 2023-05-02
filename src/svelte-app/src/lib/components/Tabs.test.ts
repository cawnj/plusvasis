import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import Tabs from './Tabs.svelte';
import { fa0 } from '@fortawesome/free-solid-svg-icons';
import html from 'svelte-htm';
import type { ComponentProps } from 'svelte';

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

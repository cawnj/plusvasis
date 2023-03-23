import type { ComponentType } from 'svelte';

export type Tab = {
	name: string;
	component: ComponentType;
};

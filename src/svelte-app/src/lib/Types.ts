import type { ComponentType } from 'svelte';

export type Job = {
	containerName: string;
	dockerImage: string;
	user: string | null;
};

export type Tab = {
	name: string;
	component: ComponentType;
};

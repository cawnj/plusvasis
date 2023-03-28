import type { ComponentType } from 'svelte';

export type Job = {
	containerName: string;
	dockerImage: string;
	user: string | null;
	shell: string;
};

export type Tab = {
	name: string;
	component: ComponentType;
};

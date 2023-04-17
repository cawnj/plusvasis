import type { Job } from '$lib/Types';
import { writable } from 'svelte/store';

export const currJobId = writable('');
export const currJob = writable({} as Job);
export const shell = writable('');
export const alloc = writable('');
export const task = writable('');
export const currJobStopped = writable(false);

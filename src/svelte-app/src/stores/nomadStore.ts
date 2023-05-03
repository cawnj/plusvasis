import { writable } from 'svelte/store';

import type { Job } from '$lib/Types';

export const currJobId = writable('');
export const currJob = writable({} as Job);
export const currJobStopped = writable(false);

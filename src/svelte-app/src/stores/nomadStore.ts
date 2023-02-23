import { writable } from 'svelte/store';

export const nomadAllocExecEndpoint = 'wss://nomad.local.cawnj.dev/v1/client/allocation/';
export const nomadAllocExecQueryParams =
	'/exec?task=server&tty=true&ws_handshake=true&command=%5B%22%2Fbin%2Fbash%22%5D';
export const job = writable('');

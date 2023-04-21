import { describe, it, expect, vi } from 'vitest';
import { fetchJobCreate } from './NomadClient';
import type { Job } from './Types';

describe('NomadClient', () => {
	it('should create a new job successfully', async () => {
		const job: Job = {
			user: 'testUser',
			containerName: 'testContainer',
			dockerImage: 'testImage',
			shell: 'testShell',
			volumes: 'testVolumes',
			env: 'testEnv',
			port: 1234,
			expose: true
		};

		const consoleSpy = vi.spyOn(console, 'log');
		await fetchJobCreate(job);
		expect(consoleSpy).toHaveBeenCalledWith('Container Created');
	});
});

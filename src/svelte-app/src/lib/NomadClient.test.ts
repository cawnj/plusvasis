import { describe, it, expect, vi } from 'vitest';
import {
	fetchJob,
	fetchJobCreate,
	fetchJobDelete,
	fetchJobIdAllocations,
	fetchJobRestart,
	fetchJobStart,
	fetchJobStop,
	fetchJobUpdate
} from './NomadClient';
import type { Job } from './Types';
import { currJobId } from '../stores/nomadStore';

describe('when we do not return an object', () => {
	currJobId.set('nomadClient');

	it('should create a new job successfully', async () => {
		const job: Job = {
			ID: 'job123',
			user: 'testUser',
			containerName: 'testContainer',
			dockerImage: 'testImage',
			shell: 'testShell',
			volumes: [['testVolumes', 'testVolumes']],
			env: [['testEnv', 'testEnv']],
			port: 1234,
			expose: true,
			cpu: 100,
			memory: 300
		};

		const consoleSpy = vi.spyOn(console, 'log');
		await fetchJobCreate(job);
		expect(consoleSpy).toHaveBeenCalledWith('Container Created');
	});

	it('should update a job successfully', async () => {
		const job: Job = {
			user: 'testUser',
			containerName: 'testContainer',
			dockerImage: 'testImage',
			shell: 'testShell',
			volumes: [['testVolumes', 'testVolumes']],
			env: [['testEnv', 'testEnv']],
			port: 1234,
			expose: true,
			cpu: 100,
			memory: 300
		};

		const consoleSpy = vi.spyOn(console, 'log');
		await fetchJobUpdate(job);
		expect(consoleSpy).toHaveBeenCalledWith('Container Updated');
	});

	it('should stop a job successfully', async () => {
		const consoleSpy = vi.spyOn(console, 'log');
		await fetchJobStop();
		expect(consoleSpy).toHaveBeenCalledWith('Container Stopped');
	});

	it('should delete a job successfully', async () => {
		const consoleSpy = vi.spyOn(console, 'log');
		await fetchJobDelete();
		expect(consoleSpy).toHaveBeenCalledWith('Container Deleted');
	});

	it('should restart a job successfully', async () => {
		const consoleSpy = vi.spyOn(console, 'log');
		await fetchJobRestart();
		expect(consoleSpy).toHaveBeenCalledWith('Container Restarted');
	});

	it('should start a job successfully', async () => {
		const consoleSpy = vi.spyOn(console, 'log');
		await fetchJobStart();
		expect(consoleSpy).toHaveBeenCalledWith('Container Started');
	});
});

describe('when we return an object', () => {
	it('should fetch a job successfully', async () => {
		const consoleSpy = vi.spyOn(console, 'log');
		const job = await fetchJob('nomadClient');
		expect(job).toBeTruthy();
		expect(consoleSpy).not.toHaveBeenCalledWith('Error');
	});

	it('should fetch job allocations successfully', async () => {
		const allocations = await fetchJobIdAllocations();
		expect(allocations).toBeTruthy();
	});
});

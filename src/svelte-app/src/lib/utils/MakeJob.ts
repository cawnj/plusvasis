import type { DockerCompose, Job } from '$lib/types/Types';

import { user } from '../../stores/auth';

let uid: string | undefined;
user.subscribe((value) => {
	uid = value?.uid;
});
export function MakeJob(formData: FormData) {
	const containerName = formData.get('containerName') as string;
	const dockerImage = formData.get('dockerImage') as string;
	const shell = formData.get('shell') as string;
	const volumeStr = formData.get('volumes') as string;
	const envStr = formData.get('env') as string;
	const port = formData.get('port') as string;
	const exposeStr = formData.get('expose') as string;
	const cpu = formData.get('cpu') as string;
	const memory = formData.get('memory') as string;

	const volumes: [string, string][] = [];
	for (const volume of volumeStr.split(',')) {
		if (volume === '') {
			continue;
		}
		volumes.push(volume.split(':') as [string, string]);
	}

	const envs: [string, string][] = [];
	for (const env of envStr.split(',')) {
		if (env === '') {
			continue;
		}
		envs.push(env.split('=') as [string, string]);
	}

	const expose: boolean = exposeStr != null;

	// When creating a job, we need to make sure we have the user's uid
	if (!uid) throw new Error('uid is undefined');

	const job: Job = {
		user: uid,
		containerName: containerName,
		dockerImage: dockerImage,
		shell: shell,
		volumes: volumes,
		env: envs,
		port: Number(port),
		expose: expose,
		cpu: Number(cpu),
		memory: Number(memory)
	};

	return job;
}

export function MakeJobsFromCompose(dockerCompose: DockerCompose): Job[] {
	const jobs: Job[] = [];
	if (!uid) throw new Error('uid is undefined');

	for (const serviceName in dockerCompose.services) {
		const service = dockerCompose.services[serviceName];

		// required
		const containerName = service.container_name;
		const dockerImage = service.image;

		// optional
		const volumes: [string, string][] = [];
		const envs: [string, string][] = [];
		let port = 0;
		let expose = false;

		// defaults
		let shell = '/bin/sh';
		let cpu = 100;
		let memory = 300;

		if (service.labels) {
			let labels = service.labels;

			// if labels string array, convert to record
			if (Array.isArray(labels)) {
				const labelsRecord: Record<string, string> = {};
				for (const label of labels) {
					const [key, value] = label.split('=');
					labelsRecord[key] = value;
				}
				labels = labelsRecord;
			}

			shell = labels['shell'] || '/bin/sh';
			cpu = Number(labels['cpu']) || 100;
			memory = Number(labels['memory']) || 300;
		}
		if (service.expose) {
			port = service.expose[0] || 0;
			expose = true;
		}
		if (service.environment) {
			let env = service.environment;

			// if env string array, convert to record
			if (Array.isArray(env)) {
				const envRecord: Record<string, string> = {};
				for (const e of env) {
					const [key, value] = e.split('=');
					envRecord[key] = value;
				}
				env = envRecord;
			}
			for (const [key, value] of Object.entries(env)) {
				envs.push([key, value]);
			}
		}
		if (service.volumes) {
			for (const vol of service.volumes) {
				const [host, container] = vol.split(':');
				volumes.push([host, container]);
			}
		}

		const job = {
			user: uid,
			containerName: containerName,
			dockerImage: dockerImage,
			shell: shell,
			volumes: volumes,
			env: envs,
			port: port,
			expose: expose,
			cpu: cpu,
			memory: memory
		};

		jobs.push(job);
	}

	return jobs;
}

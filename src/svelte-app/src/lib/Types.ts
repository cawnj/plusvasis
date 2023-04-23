import type { IconDefinition } from '@fortawesome/free-solid-svg-icons';
import type { ComponentType } from 'svelte';

export type Job = {
	user: string | null;
	containerName: string;
	dockerImage: string;
	shell: string;
	volumes: [string, string][];
	env: [string, string][];
	port: number;
	expose: boolean;
	[key: string]: string | null | number | [string, string][] | boolean;
};

type JobField = {
	title: string;
	type: string;
	options?: { value: string; name: string }[];
	placeholder: string | boolean;
	info: string;
};

export const JobFields: { key: string; value: JobField }[] = [
	{
		key: 'containerName',
		value: {
			title: 'Container Name',
			type: 'input',
			placeholder: 'alpine',
			info: 'This is used to identify the container.'
		}
	},
	{
		key: 'dockerImage',
		value: {
			title: 'Docker Image',
			type: 'input',
			placeholder: 'alpine:latest',
			info: 'Images will be pulled from a registry, e.g. Docker Hub.'
		}
	},
	{
		key: 'shell',
		value: {
			title: 'Shell',
			type: 'option',
			options: [
				{ value: '/bin/sh', name: '/bin/sh' },
				{ value: '/bin/bash', name: '/bin/bash' },
				{ value: '/bin/zsh', name: '/bin/zsh' }
			],
			placeholder: '/bin/sh',
			info: 'The shell to use when executing commands within the container.'
		}
	},
	{
		key: 'volumes',
		value: {
			title: 'Volumes',
			type: 'input',
			placeholder: 'docker_volume:/mnt/volume',
			info: 'For persistant storage, volumes are required. You can add multiple volumes by separating them with a comma.'
		}
	},
	{
		key: 'env',
		value: {
			title: 'Environment Variables',
			type: 'input',
			placeholder: 'ENV_VAR=VALUE',
			info: 'These will be passed to the container. You can add multiple environment variables by separating them with a comma.\nRefer to other containers with the {{container_name}} syntax - e.g. DB_URL=postgres://{{db}}'
		}
	},
	{
		key: 'port',
		value: {
			title: 'Port',
			type: 'input',
			placeholder: '8080',
			info: 'The port to expose to the host.'
		}
	},
	{
		key: 'expose',
		value: {
			title: 'Expose',
			type: 'toggle',
			placeholder: true,
			info: 'Enabling this setting will make your container publicly accessible.'
		}
	}
];

export type Tab = {
	name: string;
	component: ComponentType;
	icon: IconDefinition;
};

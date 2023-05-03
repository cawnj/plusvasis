// source: https://github.com/davipon/svelte-component-test-recipes/blob/main/vite.config.ts

import { sveltekit } from '@sveltejs/kit/vite';
import type { UserConfig } from 'vite';
import { configDefaults, type UserConfig as VitestConfig } from 'vitest/config';

const config: UserConfig & { test: VitestConfig['test'] } = {
	plugins: [sveltekit()],
	define: {
		'import.meta.vitest': 'undefined'
	},
	resolve: {
		alias: {
			clsx: 'clsx/dist/clsx.js'
		}
	},
	test: {
		globals: true,
		environment: 'jsdom',
		includeSource: ['src/**/*.{js,ts,svelte}'],
		setupFiles: ['./setupTest.ts', './src/mocks/setup.ts'],
		coverage: {
			exclude: ['setupTest.ts', 'src/mocks']
		},
		deps: {
			inline: [/msw/]
		},
		// Exclude playwright tests folder
		exclude: [...configDefaults.exclude, 'tests']
	},
	optimizeDeps: {
		exclude: [
			'codemirror',
			'@codemirror/language',
			'@codemirror/leagcy-modes',
			'@codemirror/theme-one-dark'
		]
	}
};

export default config;

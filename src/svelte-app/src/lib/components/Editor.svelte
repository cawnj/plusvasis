<script lang="ts">
	import CodeMirror from 'svelte-codemirror-editor';
	import { oneDark } from '@codemirror/theme-one-dark';
	import * as yamlMode from '@codemirror/legacy-modes/mode/yaml';
	import { StreamLanguage, LanguageSupport } from '@codemirror/language';

	let value: string = '';
	let yaml: LanguageSupport = new LanguageSupport(StreamLanguage.define(yamlMode.yaml));
</script>

<CodeMirror
	bind:value
	lang={yaml}
	theme={oneDark}
	styles={{
		'&': {
			width: '100%',
			height: '40rem'
		}
	}}
/>

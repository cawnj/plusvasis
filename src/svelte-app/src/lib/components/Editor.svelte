<script lang="ts">
	import { LanguageSupport, StreamLanguage } from '@codemirror/language';
	import * as yamlMode from '@codemirror/legacy-modes/mode/yaml';
	import { oneDark } from '@codemirror/theme-one-dark';
	import CodeMirror from 'svelte-codemirror-editor';

	export let value = '';
	let yaml: LanguageSupport = new LanguageSupport(StreamLanguage.define(yamlMode.yaml));
</script>

<CodeMirror
	bind:value
	lang={yaml}
	theme={oneDark}
	styles={{
		'&': {
			width: '100%',
			height: '70vh'
		}
	}}
/>

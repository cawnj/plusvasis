<script lang="ts">
	import { LanguageSupport, StreamLanguage } from '@codemirror/language';
	import * as yamlMode from '@codemirror/legacy-modes/mode/yaml';
	import { oneDark } from '@codemirror/theme-one-dark';
	import CodeMirror from 'svelte-codemirror-editor';

	export let value = '';
	const yaml: LanguageSupport = new LanguageSupport(StreamLanguage.define(yamlMode.yaml));

	const placeholder: string = `# example docker-compose
# with custom plusvasis labels

version: "3.8"

services:
  nginx:
    container_name: "nginx"
    image: "nginx"
    expose:
      - 80
    environment:
      TEST_VAR: "Hello World!"
    volumes:
      - "data:/usr/share/nginx"
    labels:
      shell: "/bin/bash"
      cpu: "100"
      memory: "300"

volumes:
  data:

`;
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
	{placeholder}
/>

<script lang="ts">
	import { onMount } from 'svelte';
	import { Button, Dropdown, Chevron, Radio } from 'flowbite-svelte';
	import { decode, getStream } from './StreamLogs';
	import { currJobStopped } from '../stores/nomadStore';

	let isStopped: boolean;
	currJobStopped.subscribe((value) => {
		isStopped = value;
	});

	let type = 'stdout';
	let logs = '';

	let preEl: HTMLPreElement;
	let stream: ReadableStream;

	// https://github.com/hashicorp/nomad/blob/main/ui/app/utils/classes/stream-logger.js
	async function readLogs() {
		logs = '';
		let streamClosed = false;
		let buffer = '';
		const decoder = new TextDecoder();
		const reader = stream.getReader();

		while (!streamClosed) {
			const { done, value } = await reader.read();
			streamClosed = done;
			buffer += decoder.decode(value, { stream: true });
			if (buffer.indexOf('}') !== -1) {
				const match = buffer.match(/(.*\})(.*)$/);
				if (match) {
					const [, chunk, newBuffer] = match;
					buffer = newBuffer;
					const result = decode(chunk);
					if (result) {
						logs += result.message;
					}
				}
			}
		}
	}

	async function changeType() {
		if (stream) stream.cancel();
		if (!isStopped) {
			stream = await getStream(type);
			readLogs();
		}
	}

	onMount(async () => {
		if (!isStopped) {
			stream = await getStream(type);
			readLogs();
		}
	});

	$: {
		if (preEl) preEl.scrollTop = preEl.scrollHeight;
	}
</script>

<div class="pb-4">
	<Button><Chevron>{type}</Chevron></Button>
	<Dropdown class="w-32 p-3 space-y-3 text-sm">
		<li>
			<Radio name="type" bind:group={type} value={'stdout'} on:change={changeType}>stdout</Radio>
		</li>
		<li>
			<Radio name="type" bind:group={type} value={'stderr'} on:change={changeType}>stderr</Radio>
		</li>
	</Dropdown>
</div>

<pre class="pre-logs" bind:this={preEl}>{logs}
</pre>

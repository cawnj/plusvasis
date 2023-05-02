<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import { Button, Dropdown, Chevron, Radio } from 'flowbite-svelte';
	import { decode, getStream } from '$lib/utils/StreamLogs';
	import { currJobStopped } from '../../stores/nomadStore';

	let isStopped: boolean;
	currJobStopped.subscribe((value) => {
		isStopped = value;
	});

	let type = 'stdout';
	let logs = '';
	let abortController: AbortController;

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
			try {
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
			} catch (error) {
				if (!(error instanceof DOMException && error.name === 'AbortError')) {
					console.error('Error reading logs:', error);
				}
				break;
			}
		}
	}

	async function changeType() {
		if (stream) abortController.abort();
		if (!isStopped) {
			abortController = new AbortController();
			stream = await getStream(type, abortController);
			readLogs();
		}
	}
	function scrollToBottom() {
		if (preEl) {
			preEl.scrollTop = preEl.scrollHeight;
		}
	}

	onMount(async () => {
		if (!isStopped) {
			abortController = new AbortController();
			stream = await getStream(type, abortController);
			readLogs();
		}
	});
	onDestroy(() => {
		if (stream) abortController.abort();
	});

	$: {
		if (logs)
			requestAnimationFrame(() => {
				setTimeout(scrollToBottom, 0);
			});
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

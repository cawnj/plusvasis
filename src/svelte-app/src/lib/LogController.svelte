<script lang="ts">
	import { alloc, task } from '../stores/nomadStore';
	import { onMount, afterUpdate } from 'svelte';
	import { b64decode } from './Base64Util';
	import type { ReadableStreamDefaultReader } from 'web-streams-polyfill/ponyfill';
	import { Button, Dropdown, Chevron, Radio } from 'flowbite-svelte';

	let allocId: string;
	let taskName: string;
	let type = 'stdout';

	alloc.subscribe((value) => {
		allocId = value;
	});
	task.subscribe((value) => {
		taskName = value;
	});

	let logs = '';
	let preEl: HTMLPreElement;
	let reader: ReadableStreamDefaultReader<Uint8Array>;

	// https://github.com/hashicorp/nomad/blob/main/ui/app/utils/stream-frames.js
	function decode(chunk: string): { offset: number; message: string } | null {
		const lines = chunk.replace(/\}\{/g, '}\n{').split('\n').filter(Boolean);
		const frames = lines.map((line) => JSON.parse(line)).filter((frame) => frame.Data);

		if (frames.length) {
			frames.forEach((frame) => (frame.Data = b64decode(frame.Data)));
			return {
				offset: frames[frames.length - 1].Offset,
				message: frames.map((frame) => frame.Data).join('')
			};
		}

		return null;
	}

	const fetchLogs = async () => {
		const urlBuilder = new URL(`https://nomad.local.cawnj.dev/v1/client/fs/logs/${allocId}`);
		urlBuilder.searchParams.append('task', taskName);
		urlBuilder.searchParams.append('type', type);
		urlBuilder.searchParams.append('follow', 'true');
		urlBuilder.searchParams.append('offset', '50000');
		urlBuilder.searchParams.append('origin', 'end');
		const url = urlBuilder.toString();

		const logFetch = async (url: string) => {
			const response = await fetch(url);
			if (response.ok) {
				return response;
			} else {
				throw new Error(response.statusText);
			}
		};

		const readerResponse = await logFetch(url);
		if (!readerResponse.body) {
			throw new Error('No response body');
		}

		// https://github.com/hashicorp/nomad/blob/main/ui/app/utils/classes/stream-logger.js
		reader = readerResponse.body.getReader();
		let streamClosed = false;
		let buffer = '';
		const decoder = new TextDecoder();

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
	};

	const handleChange = async () => {
		logs = '';
		if (reader) {
			reader.cancel();
		}
		await fetchLogs();
	};

	onMount(() => {
		fetchLogs();
	});

	afterUpdate(() => {
		if (preEl) {
			preEl.scrollTop = preEl.scrollHeight;
		}
	});
</script>

<div class="pb-4">
	<Button><Chevron>{type}</Chevron></Button>
	<Dropdown class="w-32 p-3 space-y-3 text-sm">
		<li>
			<Radio name="type" bind:group={type} value={'stdout'} on:change={handleChange}>stdout</Radio>
		</li>
		<li>
			<Radio name="type" bind:group={type} value={'stderr'} on:change={handleChange}>stderr</Radio>
		</li>
	</Dropdown>
</div>

<pre class="pre-logs" bind:this={preEl}>{logs}
</pre>

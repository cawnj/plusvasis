<script lang="ts">
	import { job, alloc, task } from '../stores/nomadStore';
	import { onMount, afterUpdate } from 'svelte';
	import { b64decode } from './Base64Util';

	let jobId: string;
	let allocId: string;
	let taskName: string;
	let type: string = 'stdout';

	job.subscribe((value) => {
		jobId = value;
	});
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
		let endOffset = 0;

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
						const { offset, message } = result;
						endOffset = offset;
						logs += message;
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

<div class="relative inline-flex pb-4">
	<select
		class="pl-2 pr-4 py-2 w-full h-full bg-gray-800 border border-gray-600 rounded-md text-white"
		bind:value={type}
		on:change={handleChange}
	>
		<option value="stdout">stdout</option>
		<option value="stderr">stderr</option>
	</select>
</div>

<pre
	class="bg-gray-900 text-white p-4 whitespace-pre-wrap max-h-96 overflow-auto scrollbar-hide"
	bind:this={preEl}>
	{logs}
</pre>

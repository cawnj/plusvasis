<script lang="ts">
	import { job, alloc, task } from '../stores/nomadStore';
	import { onMount } from 'svelte';
	import { b64decode } from './Base64Util';

	export let jobId = '';
	export let allocId = '';
	export let taskName = '';

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
	let reader: ReadableStreamReader<Uint8Array> | null = null;

	const logFetch = async (url: string) => {
		const response = await fetch(url);
		if (response.ok) {
			return response;
		} else {
			throw new Error(response.statusText);
		}
	};

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
		urlBuilder.searchParams.append('type', 'stdout');
		urlBuilder.searchParams.append('follow', 'true');
		urlBuilder.searchParams.append('offset', '50000');
		urlBuilder.searchParams.append('origin', 'end');
		const url = urlBuilder.toString();

		const readerResponse = await logFetch(url);
		if (!readerResponse.body) {
			throw new Error('No response body');
		}
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

	onMount(() => {
		fetchLogs();
	});
</script>

<pre class="text-white">{logs}</pre>

<style>
	pre {
		background-color: #0d1117;
		color: white;
		padding: 1rem;
		font-family: monospace;
		white-space: pre-wrap;
		word-wrap: break-word;
	}
</style>

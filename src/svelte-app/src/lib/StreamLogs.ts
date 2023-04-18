import { currJob, currJobId } from '../stores/nomadStore'
import { hostname } from '../stores/environmentStore'
import type { Job } from './Types';
import { b64decode } from './Base64Util';

let job: Job
let jobId: string
currJob.subscribe((value) => {
    job = value
})
currJobId.subscribe((value) => {
    jobId = value
})

// https://github.com/hashicorp/nomad/blob/main/ui/app/utils/stream-frames.js
export function decode(chunk: string): { offset: number; message: string } | null {
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

export async function getReader(type: string) {
    const urlBuilder = new URL(`${hostname}/job/${jobId}/logs`);
    urlBuilder.searchParams.append('task', job.containerName);
    urlBuilder.searchParams.append('type', type);
    urlBuilder.searchParams.append('follow', 'true');
    urlBuilder.searchParams.append('offset', '50000');
    urlBuilder.searchParams.append('origin', 'end');
    const url = urlBuilder.toString();

    const response = await fetch(url);
    if (!response.ok) {
        throw new Error(response.statusText);
    } else if (!response.body) {
        throw new Error('No response body');
    }

    return response.body.getReader();
}

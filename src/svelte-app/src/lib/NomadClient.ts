import { hostname } from "../stores/environmentStore";
import { currJobId } from "../stores/nomadStore";
import type { Job } from "$lib/Types";

let jobId: string;
currJobId.subscribe((value) => {
    jobId = value;
});

export async function fetchJobCreate(job: Job) {
    const url = `${hostname}/jobs`;
    const res = await fetch(url, {
        method: 'POST',
        body: JSON.stringify(job),
        headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`
        }
    });

    if (res.ok) {
        console.log('Container Created');
    } else {
        console.log('Error');
    }
}

export async function fetchJobUpdate(job: Job) {
    const url = `${hostname}/job/${jobId}`;
    const res = await fetch(url, {
        method: 'POST',
        body: JSON.stringify(job),
        headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`
        }
    });

    if (res.ok) {
        console.log('Container Updated');
    } else {
        console.log('Error');
    }
}

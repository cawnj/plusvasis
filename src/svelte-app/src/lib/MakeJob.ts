import type { Job } from "$lib/Types";

export function MakeJob(formData: FormData){
    const containerName = formData.get('containerName') as string;
    const dockerImage = formData.get('dockerImage') as string;
    const shell = formData.get('shell') as string;
    const volumeStr = formData.get('volumes') as string;
    const envStr = formData.get('env') as string;
    const port = formData.get('port') as string;
    const exposeStr = formData.get('expose') as string;

    const volumes: [string, string][] = [];
    for (const volume of volumeStr.split(',')) {
        if (volume === '') {
            continue;
        }
        volumes.push(volume.split(':') as [string, string]);
    }

    const envs: [string, string][] = [];
    for (const env of envStr.split(',')) {
        if (env === '') {
            continue;
        }
        envs.push(env.split('=') as [string, string]);
    }

    const expose: boolean = exposeStr != null;

    const job: Job = {
        user: localStorage.getItem('uid'),
        containerName: containerName,
        dockerImage: dockerImage,
        shell: shell,
        volumes: volumes,
        env: envs,
        port: Number(port),
        expose: expose,
    };

    return job;
};
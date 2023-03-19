<script lang="ts">
	import { goto } from '$app/navigation';
	import Nav from '$lib/NavBar.svelte';
	import NomadController from '$lib/NomadController.svelte';
	import xtermIcon from '$lib/assets/xTerm.png';
	import { onMount } from 'svelte';
	import { hostname } from '../stores/environmentStore';

	let nomadControllerComponent: NomadController;

	let jobs: any[] = [];
	onMount(async () => {
		const res = await fetch(`${hostname}/jobs`);
		jobs = await res.json();
	});
</script>

<Nav />
<button class="mb-4 btn btn-blue" on:click={() => goto('/create')}>Create Container</button>
{#each jobs as job}
	{#if job.Meta && job.Meta.user === localStorage.getItem('uid')}
		<ul>
			<div class="mt-3">
				<div class="div-container">
					<a href="/container/{job.ID}">
						<div class="card-body login-form">
							<div class="flex items-center">
								<img alt="The project logo" src={xtermIcon} class="mr-3 h-6 sm:h-14 float-left" />
								<h5 class="h5">{job.Name}</h5>
							</div>
						</div>
					</a>
				</div>
			</div>
		</ul>
	{/if}
{/each}

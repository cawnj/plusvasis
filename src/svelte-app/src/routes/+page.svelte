<script lang="ts">
	import { goto } from '$app/navigation';
	import Nav from '$lib/NavBar.svelte';
	import NomadController from '$lib/NomadController.svelte';
	import xtermIcon from '$lib/assets/xTerm.png';

	let getContainerClicked: boolean;
	let getContainerCreatedClicked: boolean;
	let containerRunning: boolean;

	export let data;
	const { jobs } = data;

	let nomadControllerComponent: NomadController;
</script>

<Nav />
<button class="mb-4 btn btn-blue" on:click={nomadControllerComponent.getContainers}
	>Get Containers</button
>
<button class="mb-4 btn btn-blue" on:click={nomadControllerComponent.getContainerCreated}
	>Create Container</button
>
{#if getContainerClicked}
	{#each nomadControllerComponent.parseData(jobs) as job}
		<ul>
			<!-- <button
				type="button"
				class="btn-purple"
				on:click={() => nomadControllerComponent.fetchJobIdAllocations(job.ID)}>{job.ID}</button
			> -->
			<div class="mt-3">
				<div class="div-container">
					<a href="/container/{job.ID}">
						<div class="card-body login-form">
							<div class="flex items-center">
								<img alt="The project logo" src={xtermIcon} class="mr-3 h-6 sm:h-14 float-left" />
								<h5 class="h5">{job.ID}</h5>
							</div>
						</div>
					</a>
				</div>
			</div>
		</ul>
	{/each}
{/if}
{#if getContainerCreatedClicked}
	{goto('/create')}
{/if}
<!-- disabled for the moment, breaks things -->
<!-- {#if containerRunning}
	<button
		type="button"
		class="btn-red"
		on:click={() => nomadControllerComponent.fetchJobIdDelete(job)}>Stop Container {job}</button
	>
{/if} -->
<NomadController
	bind:this={nomadControllerComponent}
	bind:getContainerClicked
	bind:getContainerCreatedClicked
	bind:containerRunning
/>

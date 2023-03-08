<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import Nav from '$lib/NavBar.svelte';
	import NomadController from '$lib/NomadController.svelte';

	let nomadControllerComponent: NomadController;
</script>

<Nav />
<h1 class="mb-4 text-4xl font-bold font-sans text-white">{$page.params.id}</h1>
<button
	type="button"
	class="btn-purple"
	on:click={() => nomadControllerComponent.fetchJobIdAllocations($page.params.id)}
	>Start Container</button
>
<button
	type="button"
	class="btn-purple"
	on:click={() => goto('/container/update/' + $page.params.id)}>Update Container</button
>
<NomadController bind:this={nomadControllerComponent} />

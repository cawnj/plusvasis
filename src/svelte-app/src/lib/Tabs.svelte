<script lang="ts">
	import type { Tab } from '$lib/Types';

	export let tabs: Tab[];
	export let activeTab = 0;

	const handleClicked = (index: number) => () => {
		activeTab = index;
	};
</script>

<div>
	<ul class="flex">
		{#each tabs as tab, index}
			<li
				class="text-sm font-semibold text-white bg-gray-800 rounded-t-lg"
				class:active={index === activeTab}
			>
				<button class="px-6 py-3 bg-gray-800 hover:bg-gray-600" on:click={handleClicked(index)}
					>{tab.name}</button
				>
			</li>
		{/each}
	</ul>
	{#each tabs as tab, index}
		{#if index === activeTab}
			<div class="p-4 bg-gray-800 rounded-b-lg">
				<svelte:component this={tab.component} />
			</div>
		{/if}
	{/each}
</div>

<style>
	li.active > button {
		color: rgb(31 41 55);
		background-color: white;
	}
</style>

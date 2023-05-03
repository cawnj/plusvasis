<script lang="ts">
	import { TabItem, Tabs } from 'flowbite-svelte';
	import Fa from 'svelte-fa';

	import type { Tab } from '$lib/types/Types';

	export let tabs: Tab[];
</script>

<Tabs style="underline">
	{#if !tabs || tabs.length === 0}
		<TabItem open>
			<div slot="title">No tabs</div>
		</TabItem>
	{:else}
		{#each tabs as tab, index}
			<TabItem open={index === 0}>
				<div slot="title" class="flex items-center gap-2 text-xs md:text-sm">
					<Fa icon={tab.icon} color="white" />
					{tab.name}
				</div>
				<svelte:component this={tab.component} />
			</TabItem>
		{/each}
	{/if}
</Tabs>

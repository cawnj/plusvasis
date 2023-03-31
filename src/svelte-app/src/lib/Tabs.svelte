<script lang="ts">
	import type { Tab } from '$lib/Types';
	import { Tabs, TabItem } from 'flowbite-svelte';
	import Fa from 'svelte-fa';

	export let tabs: Tab[];
</script>

<Tabs style="underline">
	{#each tabs as tab, index}
		<TabItem open={index === 0}>
			<div slot="title" class="flex items-center gap-2">
				<Fa icon={tab.icon} color="white" />
				{tab.name}
			</div>
			<svelte:component this={tab.component} />
		</TabItem>
	{/each}
</Tabs>

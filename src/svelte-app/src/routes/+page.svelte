<script lang="ts">
	import Nav from '$lib/components/NavBar.svelte';
	import Index from '$lib/components/IndexPage.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<Nav />
<Index {data} />

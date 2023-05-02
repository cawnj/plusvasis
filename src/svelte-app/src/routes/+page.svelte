<script lang="ts">
	import Nav from '$lib/NavBar.svelte';
	import Index from '$lib/IndexPage.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<Nav />
<Index {data} />

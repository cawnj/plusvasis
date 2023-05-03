<script lang="ts">
	import Index from '$lib/components/IndexPage.svelte';
	import Nav from '$lib/components/NavBar.svelte';

	import type { PageData } from './$types';

	export let data: PageData;
</script>

<Nav />
<Index {data} />

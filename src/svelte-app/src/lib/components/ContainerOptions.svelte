<script lang="ts">
	import { faPlay, faRefresh, faStop, faTrash } from '@fortawesome/free-solid-svg-icons';
	import { Button, ButtonGroup } from 'flowbite-svelte';
	import Fa from 'svelte-fa';

	import { goto } from '$app/navigation';
	import {
		fetchJobDelete,
		fetchJobRestart,
		fetchJobStart,
		fetchJobStop
	} from '$lib/utils/NomadClient';

	import { currJobStopped } from '../../stores/nomadStore';

	let width: number;
	let size: 'xs' | 'sm' | 'md' | 'lg' | 'xl' | undefined;

	$: {
		size = width < 768 ? 'xs' : 'md';
	}
</script>

<svelte:window bind:innerWidth={width} />
<ButtonGroup {size}>
	<Button
		{size}
		disabled={!$currJobStopped}
		on:click={() => fetchJobStart().then(() => window.location.reload())}
	>
		<Fa icon={faPlay} color="green" class="md:mr-2" />
		<span class="hidden md:block">Start</span>
	</Button>
	<Button
		{size}
		disabled={$currJobStopped}
		on:click={() => fetchJobRestart().then(() => window.location.reload())}
	>
		<Fa icon={faRefresh} color="orange" class="md:mr-2" />
		<span class="hidden md:block">Restart</span>
	</Button>
	<Button
		{size}
		disabled={$currJobStopped}
		on:click={() => fetchJobStop().then(() => window.location.reload())}
	>
		<Fa icon={faStop} color="red" class="md:mr-2" />
		<span class="hidden md:block">Stop</span>
	</Button>
	<Button {size} on:click={() => fetchJobDelete().then(() => goto('/'))}>
		<Fa icon={faTrash} class="md:mr-2" />
		<span class="hidden md:block">Delete</span>
	</Button>
</ButtonGroup>

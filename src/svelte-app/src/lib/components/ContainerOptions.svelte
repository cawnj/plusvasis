<script lang="ts">
	import { currJobStopped } from '../../stores/nomadStore';
	import { Button, ButtonGroup } from 'flowbite-svelte';
	import { faPlay, faStop, faTrash, faRefresh } from '@fortawesome/free-solid-svg-icons';
	import { fetchJobStop, fetchJobDelete, fetchJobRestart, fetchJobStart } from '$lib/utils/NomadClient';
	import Fa from 'svelte-fa';
	import { goto } from '$app/navigation';

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

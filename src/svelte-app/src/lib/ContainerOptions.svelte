<script lang="ts">
	import { currJobStopped } from '../stores/nomadStore';
	import { Button, ButtonGroup } from 'flowbite-svelte';
	import { faPlay, faStop, faTrash, faRefresh } from '@fortawesome/free-solid-svg-icons';
	import { fetchJobStop, fetchJobDelete, fetchJobRestart, fetchJobStart } from '$lib/NomadClient';
	import Fa from 'svelte-fa';
	import { goto } from '$app/navigation';
</script>

<ButtonGroup>
	<Button
		disabled={!$currJobStopped}
		on:click={() => fetchJobStart().then(() => window.location.reload())}
	>
		<Fa icon={faPlay} color="green" class="mr-2" />
		Start
	</Button>
	<Button
		disabled={$currJobStopped}
		on:click={() => fetchJobRestart().then(() => window.location.reload())}
	>
		<Fa icon={faRefresh} color="orange" class="mr-2" />
		Restart
	</Button>
	<Button
		disabled={$currJobStopped}
		on:click={() => fetchJobStop().then(() => window.location.reload())}
	>
		<Fa icon={faStop} color="red" class="mr-2" />
		Stop
	</Button>
	<Button on:click={() => fetchJobDelete().then(() => goto('/'))}>
		<Fa icon={faTrash} class="mr-2" />
		Delete
	</Button>
</ButtonGroup>

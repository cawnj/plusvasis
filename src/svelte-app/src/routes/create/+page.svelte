<script lang="ts">
	import NomadController from '$lib/NomadController.svelte';

	let nomadControllerComponent: NomadController;
</script>

<div class="mb-3 mt-3">
	<label for="containerNameInput" class="txt-input-label">Container Name</label>
	<input
		type="containerName"
		class="txt-input"
		id="containerNameInput"
		aria-describedby="containerNameHelp"
		placeholder="Container Name"
	/>
</div>
<div class="mb-3">
	<label for="imageInput" class="txt-input-label">Docker Image</label>
	<input
		type="dockerImage"
		class="txt-input"
		id="dockerImageInput"
		aria-describedby="dockerImageNameHelp"
		placeholder="Docker Image"
	/>
</div>
<div class="mb-3">
	<label for="portsInput" class="txt-input-label">Ports</label>
	<input
		type="ports"
		class="txt-input"
		id="portsInput"
		aria-describedby="portsHelp"
		placeholder="Ports"
	/>
</div>
<button class="mb-4 btn btn-blue" on:click={() => nomadControllerComponent.fetchJobCreate()}
	>Create Container</button
>

<div hidden>
	<NomadController bind:this={nomadControllerComponent} />
</div>

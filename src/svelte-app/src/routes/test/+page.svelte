<script lang="ts">
	import { user, signOut, signIn } from '../../stores/auth';
	import { hostname } from '../../stores/environmentStore';
	import { page } from '$app/stores';
</script>

<div
	class="flex flex-col justify-center items-center text-white h-screen w-100% space-y-6 break-all px-20"
>
	{#if $user}
		<div>
			<span>{$user.displayName}</span>
		</div>
		<div>
			<span>{$page.data.uid}</span>
		</div>
		<div>
			<span>{$page.data.token}</span>
		</div>
		<button on:click={() => signOut()}>Log out</button>
	{:else}
		<span>Not logged in.</span>
		<button on:click={() => signIn()}>Log in</button>
	{/if}
</div>

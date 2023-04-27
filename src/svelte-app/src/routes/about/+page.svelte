<script>
	import Nav from '$lib/NavBar.svelte';
	import Fa from 'svelte-fa';
	import {
		faInfoCircle,
		faEarthAmericas,
		faNetworkWired,
		faGears,
		faBarsProgress
	} from '@fortawesome/free-solid-svg-icons';
	import { Alert } from 'flowbite-svelte';
</script>

<Nav />
<div class="px-4 md:px-16 mb-4 text-white">
	<h1 class="text-xl md:text-2xl font-bold">About</h1>
	<Alert class="my-4 max-w-md">
		<span slot="icon"><Fa icon={faInfoCircle} /> </span>
		<span class="font-medium">Here's a lil bit of info about us...</span>
	</Alert>
	<div class="flex items-center my-1">
		<Fa icon={faEarthAmericas} />
		<span class="font-medium text-white ml-2"
			>PlusVasis is an intuitive online platform that simplifies container management and
			orchestration for developers.</span
		>
	</div>
	<div class="flex items-center my-1 mt-2">
		<Fa icon={faNetworkWired} />
		<span class="font-medium text-white ml-2"
			>It offers seamless networking, security, and scalability, allowing users to spin up
			containers for their websites or databases effortlessly.</span
		>
	</div>
	<div class="flex items-center my-1 mt-2">
		<Fa icon={faGears} />
		<span class="font-medium text-white ml-2"
			>PlusVasis is powered by modern frameworks and infrastructure solutions, making it easy for
			developers to manage their containers without worrying about complex backend processes.</span
		>
	</div>
	<div class="flex items-center my-1 mt-2">
		<Fa icon={faBarsProgress} />
		<span class="font-medium text-white ml-2"
			>Overall, PlusVasis streamlines container management and enables developers to focus on
			building their applications.</span
		>
	</div>
</div>

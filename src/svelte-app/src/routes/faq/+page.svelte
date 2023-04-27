<script>
	import Nav from '$lib/NavBar.svelte';
	import Fa from 'svelte-fa';
	import { faInfoCircle, faCheck } from '@fortawesome/free-solid-svg-icons';
	import { Alert } from 'flowbite-svelte';
</script>

<Nav />
<div class="px-4 md:px-16 mb-4 text-white">
	<h1 class="text-xl md:text-2xl font-bold">FAQ</h1>
	<Alert class="my-4 max-w-md">
		<span slot="icon"><Fa icon={faInfoCircle} /> </span>
		<span class="font-medium">What is PlusVasis, and what does it do?</span>
	</Alert>
	<div class="flex items-center my-1">
		<Fa icon={faCheck} />
		<span class="font-medium text-white ml-2"
			>In simple terms, PlusVasis is a container management system. For more details please visit
			our about page.</span
		>
	</div>
	<Alert class="my-4 max-w-md">
		<span slot="icon"><Fa icon={faInfoCircle} /> </span>
		<span class="font-medium">How does PlusVasis simplify container management for developers?</span
		>
	</Alert>
	<div class="flex items-center my-1">
		<Fa icon={faCheck} />
		<span class="font-medium text-white ml-2"
			>PlusVasis simplifies container management for developers by providing a welcoming and user
			friendly UX, offering users a pleasant experience.</span
		>
	</div>
	<Alert class="my-4 max-w-md">
		<span slot="icon"><Fa icon={faInfoCircle} /> </span>
		<span class="font-medium">What kind of containers can I spin up using PlusVasis?</span>
	</Alert>
	<div class="flex items-center my-1">
		<Fa icon={faCheck} />
		<span class="font-medium text-white ml-2"
			>Users have the ability to spin up any container that is on the Docker Registry.</span
		>
	</div>
	<Alert class="my-4 max-w-md">
		<span slot="icon"><Fa icon={faInfoCircle} /> </span>
		<span class="font-medium">How does PlusVasis ensure network security for my containers?</span>
	</Alert>
	<div class="flex items-center my-1">
		<Fa icon={faCheck} />
		<span class="font-medium text-white ml-2">Security is enforced by a Traefik reverse proxy.</span
		>
	</div>
	<Alert class="my-4 max-w-md">
		<span slot="icon"><Fa icon={faInfoCircle} /> </span>
		<span class="font-medium">What scalability options does PlusVasis offer?</span>
	</Alert>
	<div class="flex items-center my-1">
		<Fa icon={faCheck} />
		<span class="font-medium text-white ml-2"
			>PlusVasis offers seamless scalability for containers. It allows users to spin up containers
			for their websites or databases effortlessly, without worrying about the complex backend
			processes.</span
		>
	</div>
	<Alert class="my-4 max-w-md">
		<span slot="icon"><Fa icon={faInfoCircle} /> </span>
		<span class="font-medium">How to use Custom DockerFile?</span>
	</Alert>
	<div class="flex items-center my-1">
		<Fa icon={faCheck} />
		<span class="font-medium text-white ml-2"
			>For a user to use a custom Dockerfile, they must first add their custom image to the Docker
			Registry. Then they can specify that custom image for usage upon the container creation page.</span
		>
	</div>
</div>

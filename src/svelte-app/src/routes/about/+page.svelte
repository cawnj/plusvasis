<script>
	import Nav from '$lib/NavBar.svelte';
	import Fa from 'svelte-fa';
	import {
		faEarthAmericas,
		faNetworkWired,
		faGears,
		faBarsProgress
	} from '@fortawesome/free-solid-svg-icons';
	import { Card, Heading, Span } from 'flowbite-svelte';
</script>

<Nav />
<Heading tag="h1" class="text-8xl font-bold text-center mb-8 text-white">What is PlusVasis?</Heading
>
<div class="px-4 md:px-16 space-y-8 text-white flex justify-center">
	<div class="grid grid-cols-2 gap-8">
		<Card horizontal>
			<div class="flex flex-col">
				<div class="flex items-center justify-center mb-4">
					<Fa icon={faEarthAmericas} class="text-4xl text-sky-400 mr-4" />
					<Heading tag="h2" class="text-2xl font-semibold">
						<Span gradient>Intuitive Platform</Span>
					</Heading>
				</div>
				<div>
					<p class="text-xl text-white">
						PlusVasis is an intuitive online platform that simplifies container management and
						orchestration for developers.
					</p>
				</div>
			</div>
		</Card>
		<Card horizontal>
			<div class="flex flex-col">
				<div class="flex items-center mb-4">
					<Fa icon={faNetworkWired} class="text-4xl text-sky-400 mr-4" />
					<Heading tag="h2" class="text-2xl font-semibold">
						<Span gradient>Seamless Networking</Span>
					</Heading>
				</div>
				<div>
					<p class="text-xl text-white">
						PlusVasis offers seamless networking, security, and scalability, allowing users to spin
						up containers for their websites or databases effortlessly.
					</p>
				</div>
			</div>
		</Card>
		<Card horizontal>
			<div class="flex flex-col">
				<div class="flex items-center mb-4">
					<Fa icon={faGears} class="text-4xl text-sky-400 mr-4" />
					<Heading tag="h2" class="text-2xl font-semibold">
						<Span gradient>Modern Infrastructure</Span>
					</Heading>
				</div>
				<div>
					<p class="text-xl text-white">
						PlusVasis is powered by modern frameworks and infrastructure solutions, making it easy
						for developers to manage their containers without worrying about complex backend
						processes.
					</p>
				</div>
			</div>
		</Card>
		<Card horizontal>
			<div class="flex flex-col">
				<div class="flex items-center mb-4">
					<Fa icon={faBarsProgress} class="text-4xl text-sky-400 mr-4" />
					<Heading tag="h2" class="text-2xl font-semibold">
						<Span gradient>Streamlined Management</Span>
					</Heading>
				</div>
				<div>
					<p class="text-xl text-white">
						PlusVasis streamlines container management and enables developers to focus on building
						their applications.
					</p>
				</div>
			</div>
		</Card>
	</div>
</div>

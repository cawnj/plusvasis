<script>
	import Nav from '$lib/components/NavBar.svelte';
	import Fa from 'svelte-fa';
	import {
		faEarthAmericas,
		faNetworkWired,
		faGears,
		faBarsProgress
	} from '@fortawesome/free-solid-svg-icons';
	import { Card, Heading, Span } from 'flowbite-svelte';

	const cards = [
		{
			icon: faEarthAmericas,
			title: 'Intuitive Platform',
			description:
				'PlusVasis is an intuitive online platform that simplifies container management and orchestration for developers.'
		},
		{
			icon: faNetworkWired,
			title: 'Seamless Networking',
			description:
				'PlusVasis offers seamless networking, security, and scalability, allowing users to spin up containers for their websites or databases effortlessly.'
		},
		{
			icon: faGears,
			title: 'Modern Infrastructure',
			description:
				'PlusVasis is powered by modern frameworks and infrastructure solutions, making it easy for developers to manage their containers without worrying about complex backend processes.'
		},
		{
			icon: faBarsProgress,
			title: 'Streamlined Management',
			description:
				'PlusVasis streamlines container management and enables developers to focus on building their applications.'
		}
	];
</script>

<Nav />
<Heading tag="h1" class="text-4xl md:text-8xl font-bold text-center mb-4 md:mb-8 text-white"
	>What is PlusVasis?</Heading
>
<div class="px-4 md:px-16 pb-4 space-y-8 text-white flex justify-center">
	<div class="grid grid-cols-2 gap-8">
		{#each cards as card}
			<Card horizontal>
				<div class="flex flex-col">
					<div class="flex items-center justify-center mb-4">
						<Fa icon={card.icon} class="text-xl md:text-4xl text-sky-400 mr-4" />
						<Heading tag="h2" class="text-lg md:text-2xl font-semibold">
							<Span gradient>{card.title}</Span>
						</Heading>
					</div>
					<div>
						<p class="text-lg md:text-xl text-white">
							{card.description}
						</p>
					</div>
				</div>
			</Card>
		{/each}
	</div>
</div>

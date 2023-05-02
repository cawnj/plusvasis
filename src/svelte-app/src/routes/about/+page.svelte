<script>
	import {
		faBarsProgress,
		faEarthAmericas,
		faGears,
		faNetworkWired
	} from '@fortawesome/free-solid-svg-icons';
	import { Card, Heading, Span } from 'flowbite-svelte';
	import Fa from 'svelte-fa';

	import Nav from '$lib/components/NavBar.svelte';

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
<Heading tag="h1" class="mb-4 text-center text-4xl font-bold text-white md:mb-8 md:text-8xl"
	>What is PlusVasis?</Heading
>
<div class="flex justify-center space-y-8 px-4 pb-4 text-white md:px-16">
	<div class="grid grid-cols-2 gap-8">
		{#each cards as card}
			<Card horizontal>
				<div class="flex flex-col">
					<div class="mb-4 flex items-center justify-center">
						<Fa icon={card.icon} class="mr-4 text-xl text-sky-400 md:text-4xl" />
						<Heading tag="h2" class="text-lg font-semibold md:text-2xl">
							<Span gradient>{card.title}</Span>
						</Heading>
					</div>
					<div>
						<p class="text-lg text-white md:text-xl">
							{card.description}
						</p>
					</div>
				</div>
			</Card>
		{/each}
	</div>
</div>

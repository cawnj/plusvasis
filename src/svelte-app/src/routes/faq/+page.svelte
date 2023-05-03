<script>
	import { faQuestionCircle } from '@fortawesome/free-solid-svg-icons';
	import { Blockquote, Heading } from 'flowbite-svelte';
	import Fa from 'svelte-fa';

	import Nav from '$lib/components/NavBar.svelte';

	const faqContent = [
		{
			question: 'What is PlusVasis, and what does it do?',
			answer:
				'In simple terms, PlusVasis is a container management system. For more details please visit our about page.'
		},
		{
			question: 'How does PlusVasis simplify container management for developers?',
			answer:
				'PlusVasis simplifies container management for developers by providing a welcoming and user friendly UX, offering users a pleasant experience.'
		},
		{
			question: 'What kind of containers can I spin up using PlusVasis?',
			answer:
				'Users have the ability to spin up any container that is available on any public Docker Registry.'
		},
		{
			question: 'How does PlusVasis ensure security for my containers?',
			answer:
				'Security is enforced by a Traefik reverse proxy and a well-integrated authentication layer using Firebase'
		},
		{
			question: 'What scalability options does PlusVasis offer?',
			answer:
				'PlusVasis offers seamless scalability for containers. It allows users to spin up containers for their websites or databases effortlessly, without worrying about the complex backend processes.'
		},
		{
			question: 'How to use a Custom Dockerfile?',
			answer:
				'For a user to use a custom Dockerfile, they must first push their custom image to a Docker Registry. Then they can specify that custom image for usage upon the container creation page.'
		}
	];
</script>

<Nav />
<Heading tag="h1" class="mb-8 text-center text-8xl font-bold text-white">FAQ</Heading>
<div class="px-4 pb-4 md:px-16">
	{#each faqContent as { question, answer }}
		<div class="my-4 space-y-1">
			<Heading tag="h1" class="flex items-center text-xl font-semibold text-white">
				<Fa icon={faQuestionCircle} class="mr-2" />
				<span>{question}</span>
			</Heading>
			<Blockquote size="lg">
				<span class="font-medium text-gray-400">{answer}</span>
			</Blockquote>
		</div>
	{/each}
</div>

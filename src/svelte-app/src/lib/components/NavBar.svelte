<script>
	import logo from '$lib/assets/logo.png';
	import { Navbar, NavBrand, NavLi, NavUl, NavHamburger } from 'flowbite-svelte';
</script>

<div class="mx-1 my-2">
	<Navbar let:hidden let:toggle>
		<NavBrand href="/">
			<img src={logo} class="mr-3 h-6 sm:h-9" alt="PlusVasis Logo" />
			<span
				class="self-center whitespace-nowrap text-xl font-bold md:text-2xl dark:text-white hover:text-blue-400"
			>
				PlusVasis
			</span>
		</NavBrand>
		<NavHamburger on:click={toggle} />
		<NavUl data-testid="navbar-ul" {hidden}>
			<NavLi href="/" nonActiveClass="hover:text-blue-400 text-xl">Home</NavLi>
			<NavLi href="/about" nonActiveClass="hover:text-blue-400 text-xl">About</NavLi>
			<NavLi href="/faq" nonActiveClass="hover:text-blue-400 text-xl">FAQ</NavLi>
			<NavLi href="/logout" nonActiveClass="hover:text-blue-400 text-xl">Sign Out</NavLi>
		</NavUl>
	</Navbar>
</div>

<script>
	import logo from '$lib/assets/logo.png';
	import { logout } from '$lib/fb';
	import {
		Navbar,
		NavBrand,
		NavLi,
		NavUl,
		NavHamburger,
		Avatar,
		Dropdown,
		DropdownItem,
		DropdownHeader,
		DropdownDivider
	} from 'flowbite-svelte';
</script>

<div class="mx-3 my-2 w-screen">
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
		<NavUl {hidden}>
			<NavLi href="/" nonActiveClass="hover:text-blue-400 text-xl">Home</NavLi>
			<NavLi href="/about" nonActiveClass="hover:text-blue-400 text-xl">About</NavLi>
			<NavLi href="/logout" nonActiveClass="hover:text-blue-400 text-xl">Sign Out</NavLi>
		</NavUl>
	</Navbar>
</div>

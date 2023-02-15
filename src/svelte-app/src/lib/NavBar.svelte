<script>
	import logo from '$lib/assets/logo.png';
</script>

<div>
	<div>
		<nav class="container px-6 py-8 mx-auto md:flex md:justify-between md:items-center">
			<div class="flex items-center justify-between">
				<img alt="The project logo" src={logo} class="mr-3 h-6 sm:h-9" />
				<a class="text-xl font-bold text-white md:text-2xl hover:text-blue-400" href="/"
					>Continens
				</a>
			</div>

			<div
				class="flex-col mt-8 space-y-4 md:flex md:space-y-0 md:flex-row md:items-center md:space-x-10 md:mt-0"
			>
				<a class="text-white hover:text-blue-400" href="/">Home</a>
				<a class="text-white hover:text-blue-400" href="/">About</a>
			</div>
		</nav>
	</div>
</div>

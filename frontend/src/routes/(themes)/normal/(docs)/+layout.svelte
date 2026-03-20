
<script lang="ts">
	import Nav from '$lib/components/docs/Nav.svelte';



	let { children } = $props();


</script>


<nav>
	<Nav/>
</nav>

<main>
	<a href="/search" class="application">URMC-HUB</a>
	{@render children?.()}
</main>

<style>

	nav {
		position: fixed;
		display: flex;
		justify-content: flex-start;
		flex-direction: column;
		height: 100%;
		width: 250px;
		top: 0;
		left: 0;
		box-sizing: border-box;
		text-align: left;
		background: var(--color-surface);
		overflow: hidden;
	}


	main {
		margin-top: -50px;
		padding-left: 270px;
	}

	a.application {
		position: fixed;
		top: 1rem;
		right: 1rem;
		color: var(--color-text);
		text-decoration: none;
	}

	a.application:hover {
		color: var(--color-text-opacity-80);
	}

	a.application:visited {
		color: var(--color-text);
	}

	
</style>

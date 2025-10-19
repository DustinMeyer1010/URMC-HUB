<script lang="ts">
	import urmc from '$lib/assets/URMC.png';
	import Nav from '../components/Nav.svelte';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={urmc} />
</svelte:head>

<Nav/>

{@render children?.()}


<style>
	:global(:root) {
		--primary: #008cff;
		--background: #212335;
		--background-surface: #2F324B;
		--text: #fdebd6;
		--complement: #ff7300;
		margin-top: 50px;
		background: var(--background);

	}

</style>


<!--
Dark Theme
----------------
Background: #212335
Background-Surface: #2F324B
Primary: #008cff
On-Primary: #fdebd6
Complementary: #ff7300
-->

<script lang="ts">

	import urmc from '$lib/assets/URMC.png';
	import Nav from '../components/Nav.svelte';
	import "../global.css"

	import { theme } from "$lib/theme";
  	import { get } from "svelte/store";

	const current = get(theme);



	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={urmc} />
</svelte:head>

<Nav/>

{@render children?.()}


<style>

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
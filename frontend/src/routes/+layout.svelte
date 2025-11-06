
<script lang="ts">

	import urmc from '$lib/assets/URMC.png';
	import "../global.css"

	import { theme } from "$lib/theme";
  	import { get } from "svelte/store";
	const current = get(theme);

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={urmc} />
</svelte:head>


{@render children?.()}

<style>

</style>

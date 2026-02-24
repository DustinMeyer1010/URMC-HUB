
<script lang="ts">

	import urmc from '$lib/assets/URMC.png';
	import "$lib/styles/global.css"

	import { theme } from "$lib/stores/theme";
  	import { get } from "svelte/store";
	get(theme);

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={urmc} />
</svelte:head>


{@render children?.()}

<style>


</style>

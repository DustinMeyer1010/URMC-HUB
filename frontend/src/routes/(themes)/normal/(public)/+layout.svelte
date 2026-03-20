<script lang="ts">
	import Nav from '$lib/components/navigation/Nav.svelte';

	let { children } = $props();
</script>

<Nav/>
{@render children?.()}

<style>
</style>

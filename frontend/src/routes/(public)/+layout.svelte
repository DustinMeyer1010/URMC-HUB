<script lang="ts">
	import Nav from '@components/Nav.svelte';

	let { children } = $props();
</script>

<Nav/>
{@render children?.()}

<style>
</style>

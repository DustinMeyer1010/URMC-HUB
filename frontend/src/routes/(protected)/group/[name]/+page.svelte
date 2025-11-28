<script lang="ts">
	import Nav from '../Nav.svelte';
	import { GroupPageStateClass, Sections } from './state.svelte';

    let {
        data
    } : {
        data: {name: string}
    } = $props()

    let PageState: GroupPageStateClass = new GroupPageStateClass()

</script>

<main>
    <h1>{data.name}</h1>
    <Nav sections={Sections} swapSection={PageState.swapSection} />
</main>



<style>
    main {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }
</style>


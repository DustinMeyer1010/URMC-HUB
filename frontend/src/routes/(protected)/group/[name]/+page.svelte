<script lang="ts">
	import Info from '$lib/components/group/Info.svelte';
	import Nav from '../Nav.svelte';
	import { GroupPageStateClass, Sections } from './state.svelte';
	import Members from '$lib/components/group/Members.svelte';
	import Add from '$lib/components/group/Add.svelte';
	import Remove from '$lib/components/group/Remove.svelte';

    let {
        data
    } : {
        data: {name: string}
    } = $props()

    let PageState: GroupPageStateClass = new GroupPageStateClass()

</script>

<main>
    <header>
        <h1>{data.name}</h1>
        <Info group={data.name}/>
        <Nav sections={Sections} swapSection={PageState.swapSection} currentSection={PageState.section} />
    </header>
    <section>
        {#if PageState.section == "MEMBERS"}
            <Members group={data.name}/>
        {:else if PageState.section == "ADD"}
            <Add group={data.name}/>
        {:else if PageState.section == "REMOVE"}
            <Remove group={data.name}/>
        {/if}
    </section>
</main>



<style>
    main {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        padding: 0;
        margin: 0;
    }

    header {
        width: 100%;
        display: flex;
        flex-direction: column;
        text-align: center;
        justify-content: center;
        position: sticky;
        padding: 10px;
        background: var(--color-bg);
        z-index: 1;
        top: 61px;
        margin-bottom: 20px;
        box-shadow: 20px 0px 10px rgba(0,0,0,0.3);
    }

    h1 {
        margin: 0;
        padding: 10px;
        font-size: 15px;
    }

    section {
        width: 100%;
        display: flex;
        justify-content: center;

    }

</style>


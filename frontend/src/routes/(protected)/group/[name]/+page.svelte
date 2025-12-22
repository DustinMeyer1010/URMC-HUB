<script lang="ts">
	import Info from '@components/Group/Info.svelte';
	import Nav from '../Nav.svelte';
	import { GroupPageStateClass, Sections } from './state.svelte';
	import Members from '@components/Group/Members.svelte';
	import Add from '@components/Group/Add.svelte';

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
        {/if}
    </section>
</main>



<style>
    main {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }

    header {
        top: 60px;
        width: 100%;
        display: flex;
        flex-direction: column;
        text-align: center;
        justify-content: center;
        position: fixed;
        padding: 0;
        background: var(--color-bg);
        z-index: 10;
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
        margin-top: 250px;
    }

</style>


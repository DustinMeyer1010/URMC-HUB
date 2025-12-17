<script lang="ts">
	import Info from '@components/Group/Info.svelte';
	import Nav from '../Nav.svelte';
	import { GroupPageStateClass, Sections } from './state.svelte';
	import Members from '@components/Group/Members.svelte';

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
        <Info group={data.name}></Info>
        <Nav sections={Sections} swapSection={PageState.swapSection} currentSection={PageState.section} />
    </header>
    <section>
        {#if PageState.section == "MEMBERS"}
            <Members group={data.name}/>
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
        top: 50px;
        width: 100%;
        display: flex;
        flex-direction: column;
        text-align: center;
        justify-content: center;
        position: fixed;
        background: var(--color-bg);
        z-index: 10;
    }

    section {
        width: 100%;
        display: flex;
        justify-content: center;
        margin-top: 300px;
    }

</style>


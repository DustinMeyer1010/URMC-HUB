<script lang="ts">
	import Info from '@components/Group/Info.svelte';
	import Nav from '../Nav.svelte';
	import { GroupPageStateClass, Sections } from './state.svelte';
	import { onMount } from 'svelte';
	import Members from '@components/Group/Members.svelte';

    let {
        data
    } : {
        data: {name: string}
    } = $props()

    let PageState: GroupPageStateClass = new GroupPageStateClass()

    onMount(async () => {
        await PageState.retrievePageData(data.name)
    })

</script>

<main>
    <h1>{data.name}</h1>
    <Nav sections={Sections} swapSection={PageState.swapSection} />

    {#if PageState.PageInfo}
        {#if PageState.section == "INFO"}
            <Info group={PageState.PageInfo}></Info>
        {:else if PageState.section == "MEMBERS"}
            <Members group={data.name}/>

        {/if}
    {/if}
</main>



<style>
    main {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }
</style>


<script lang="ts">
    
    // === COMPONENTS ===
    import Group from "./cards/Group.svelte";
    import Computer from "./cards/Computer.svelte";
    import Printer from "./cards/Printer.svelte";
	import User from "./cards/User.svelte";
	import Drive from "./cards/Drive.svelte";
	import { Search } from "$lib/types/search";
	import type { Component } from 'svelte';


    let {
        data,
        filter
    } : {
        filter: Search.Filter,
        data: Search.Results
    } = $props()

    const FilterMap: Record<string, { items: any[], Component: Component<{item: any, idx: number}> }> = {
        COMPUTERS: { items: data.computers, Component: Computer },
        PRINTERS:  { items: data.printers, Component: Printer },
        USERS:     { items: data.users, Component: User },
        GROUPS:    { items: data.groups, Component: Group },
        DRIVES:    { items: data.drives, Component: Drive },
    };
    let Card = $derived(FilterMap[filter].Component)

</script>

<div>
        {#each FilterMap[filter].items as item, idx}
            <Card item={item} idx={idx} />
        {:else}
            {@render NotFound(filter)}
        {/each}
</div>


{#snippet NotFound(filter: Search.Filter)}
    <h1>No {filter} Found</h1>
{/snippet}



<style>

    h1 {
        text-align: center;
        font-size: 15px;
    }

    div {

        display: flex;
        padding: 1rem;
        flex-direction: column;
        gap: 1rem;
    }

</style>
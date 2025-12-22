<script lang="ts">
    // === TYPES ===
    import type { Filters } from '@t/filters'
    
    // === COMPONENTS ===
    import Group from "./Cards/Group.svelte";
    import Computer from '@components/Cards/Computer.svelte'
	import Printer from "./Cards/Printer.svelte";
	import User from "./Cards/User.svelte";
	import Drive from "./Cards/Drive.svelte";
	import type { AllResults } from "@t/resutls";
	import type { Component } from 'svelte';


    let {
        data,
        filter
    } : {
        filter: Filters,
        data: AllResults
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


{#snippet NotFound(filter: Filters)}
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
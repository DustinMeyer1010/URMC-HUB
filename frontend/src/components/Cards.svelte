<script lang="ts">
    import type { ComputerSimpleInfo } from "@t/computer"
    import type { PrinterSimpleInfo } from "@t/printer"
    import type { UserSimpleInfo } from "@t/user";
    import type { Results } from "@t/filters"
    import type { Groups as GroupFilter } from '@t/filters'
    import type { GroupSimpleInfo } from "@t/group";
	import type { DriveSimpleInfo } from "@t/drive";

    import Group from "./Cards/Group.svelte";
    import Computer from '@components/Cards/Computer.svelte'
	import Printer from "./Cards/Printer.svelte";
	import User from "./Cards/User.svelte";
	import Drive from "./Cards/Drive.svelte";


    let {
        items,
        filter
    } : {
        filter: GroupFilter,
        items: Results
    } = $props()

</script>

<div>
    {#if filter === 'COMPUTERS'}
        {#each (items as ComputerSimpleInfo[]) as computer, idx}
            <Computer {computer} {idx}/>
        {:else}
            {@render NotFound(filter)} 
        {/each}
    {:else if filter === 'PRINTERS'}
        {#each (items.slice(0, 100) as PrinterSimpleInfo[]) as printer, idx}
            <Printer {printer} {idx}/>
        {:else}
            {@render NotFound(filter)}    
        {/each}
    {:else if filter === 'USERS'}
        {#each (items as UserSimpleInfo[]) as user, idx}
            <User {user} {idx}/>
        {:else}
            {@render NotFound(filter)}
        {/each}
    {:else if filter === 'GROUPS'}
        {#each (items as GroupSimpleInfo[]) as group, idx}
            <Group {group} {idx}/>
        {:else}
            {@render NotFound(filter)}
        {/each}
    {:else if filter === 'DRIVES'}
        {#each (items as DriveSimpleInfo[]) as drive, idx}
            <Drive {drive} {idx}/>
        {:else}
            {@render NotFound(filter)}
        {/each}
    {/if}
</div>

{#snippet NotFound(filter: GroupFilter)}
    <span>No {filter} Found</span>
{/snippet}



<style>
    div {

        display: flex;
        padding: 1rem;
        flex-direction: column;
        gap: 1rem;
    }

</style>
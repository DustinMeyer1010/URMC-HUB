<script lang="ts">
    import type { ComputerSimpleInfo } from "@t/computer"
    import type { PrinterSimpleInfo } from "@t/printer"
    import type { UserSimpleInfo } from "@t/user";
    import type { Results } from "@t/filters"
    import type { Groups as GroupFilter } from '@t/filters'
    import type { GroupSimpleInfo } from "@t/group";

    import Group from "./Cards/Group.svelte";
    import Computer from '@components/Cards/Computer.svelte'
	import Printer from "./Cards/Printer.svelte";
	import User from "./Cards/User.svelte";


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
        {/each}
    {:else if filter === 'PRINTERS'}
        {#each (items.slice(0, 100) as PrinterSimpleInfo[]) as printer, idx}
            <Printer {printer} {idx}/>    
        {/each}
    {:else if filter === 'USERS'}
        {#each (items as UserSimpleInfo[]) as user, idx}
            <User {user} {idx}/>
        {/each}
    {:else if filter === 'GROUPS'}
        {#each (items as GroupSimpleInfo[]) as group, idx}
            <Group {group} {idx}/>
        {/each}
    {/if}
    {#if items == null || items.length <= 0}
            No {filter} found
    {/if}
</div>



<style>
    div {
        padding: 1rem;
        display: flex;
        flex-direction: column;
        gap: 1rem;
    }

</style>
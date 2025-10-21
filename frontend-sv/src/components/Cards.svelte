<script lang="ts">
    import type { ComputerSimpleInfo } from "@types/computer"
    import type { PrinterSimpleInfo } from "@types/printer"
    import type { UserSimpleInfo } from "@types/user";
    import type { Results } from "@types/filters"
    import type { Groups as GroupFilter } from '@types/filters'
    import type { GroupSimpleInfo } from "@types/group";

    import Groups from "./Cards/Groups.svelte";
    import Computers from '@components/Cards/Computers.svelte'
	import Printers from "./Cards/Printers.svelte";
	import Users from "./Cards/Users.svelte";


    let {
        items,
        filter
    } : {
        filter: GroupFilter,
        items: Results
    } = $props()

</script>

{#if filter === 'COMPUTERS'}
    <Computers computers={items as ComputerSimpleInfo[]}/>
{:else if filter === 'PRINTERS'}
    <Printers printers={items as PrinterSimpleInfo[]}/>
{:else if filter === 'USERS'}
    <Users users={items as UserSimpleInfo[]}/>
{:else if filter === 'GROUPS'}
    <Groups groups={items as GroupSimpleInfo[]}/>
{/if}



<style>
    :global(ul) {
        opacity: 0;
        transform: translateY(20px);
        animation: slideIn 0.2s ease-out forwards;
        animation-delay: var(--delay);
}


    @keyframes slideIn {
	from {
		opacity: 0;
		transform: translateY(20px);
	}
	to {
		opacity: 1;
		transform: translateY(0);
	}
}
</style>
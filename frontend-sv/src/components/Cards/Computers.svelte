<script lang="ts">
    import type { ComputerSimpleInfo } from "@types/computer";


let {
    computers
} : {
    computers: ComputerSimpleInfo[]
} = $props()


</script>

{#if computers.length != 0}
{#each computers as c}
    <ul class:disabled={c.ou.toLowerCase().includes("disabled")}>
        <li>{c.name}</li>
        <li>{c.ou}</li>
        <li>{c.operating_system}</li>
    </ul>
{/each}
{:else}
    No computers found
{/if}



<style>
    ul {
        list-style: none;
        color: var(--text)
    }
    li {
        margin-top: 5px; 
    }

    ul.disabled {
        color: red
    }
</style>




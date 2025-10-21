<script lang="ts">
    import type { ComputerSimpleInfo } from "@types/computer";


let {
    computers
} : {
    computers: ComputerSimpleInfo[]
} = $props()


</script>

{#if computers.length != 0}
    {#each computers as c,i}
        <ul class:disabled={c.ou.toLowerCase().includes("disabled")} style="--delay: {Math.min(i * 50, 2000)}ms">
            <a href="#">test</a>
            <li>{c.name}</li>
            <li>{c.ou}</li>
            <li>{c.operating_system}</li>
        </ul>
    {/each}
{:else}
    No computers found
{/if}



<style>
    ul.disabled {
        color: red
    }
</style>




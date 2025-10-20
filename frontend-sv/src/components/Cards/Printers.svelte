<script lang="ts">
    import type { PrinterSimpleInfo } from "@types/printer";


    let {
        printers
    } : {
        printers: PrinterSimpleInfo[]
    } = $props()

</script>

{#if printers.length != 0}
{#each printers as p}
    <ul>
        <li>{p.server}</li>
        <li>{p.queue}</li>
        <li>{p.model}</li>
        <li>{p.ip}</li>
        <li>{p.print_processor}</li>
        <li>{p.location}</li>
        <li>{p.notes}</li>
    </ul>
{/each}
{:else}
No Printers Found
{/if}
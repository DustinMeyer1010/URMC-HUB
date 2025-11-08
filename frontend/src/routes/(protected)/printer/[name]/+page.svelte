<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import type { PrinterSimpleInfo } from '@t/printer';
	import { onMount } from 'svelte';


    let printer: PrinterSimpleInfo | null = $state(null)

    let {
        data
    } : {
        data: {name: string}
    } = $props()

    onMount(async () => {
        let urlParams = page.url.searchParams;
        let queue = urlParams.get("queue")

        if (!queue) {
            goto("/search")
            return
        }

        let res = await fetch(`http://localhost:8000/api/printer/info/${data.name}?queue=${queue}`)

        if (res.status != 200) {
            goto("/search")
            return
        }
        

        printer = await res.json() as PrinterSimpleInfo

        console.log(printer)
    })



</script>

{#if printer != null}
    <h1>{`\\\\${printer.server}\\${printer.queue}`}</h1>
    <a href={`http://${printer.ip.replace(/_.*/, "")}`} target="_blank">Open Printer Settings</a>
{/if}

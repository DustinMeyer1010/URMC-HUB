<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import PrinterStatus from '@components/Printer/PrinterStatus.svelte';
	import type { PrinterSimpleInfo } from '@t/printer';
	import { onMount } from 'svelte';


    let printer: PrinterSimpleInfo | null = $state(null)
    let isOnline: boolean = $state(false)
    let eRecordPrinter: boolean = $state(false)

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

        res = await fetch(`http://localhost:8000/api/printer/ping/${printer.ip.replace(/_.*/, "")}`)


        const str = await res.text();
        isOnline = str.includes("Successfully")
        eRecordPrinter = printer.server.includes("eRcd-CPA")
    })


</script>

{#if printer != null}
    <section>
        <header>
            <h1>{`\\\\${printer.server}\\${printer.queue}`}</h1>
            <sup><PrinterStatus {isOnline} {eRecordPrinter} ip={printer.ip}/></sup>
        </header>
        <div>
            <span><b>Location: </b>{printer.location}</span>
            <span>
                {#if isOnline && !eRecordPrinter}
                    <b>IP: </b><a href={`http://${printer.ip.replace(/_.*/, "")}`}>{printer.ip}</a>
                {:else}
                    <b>IP: </b>{printer.ip}
                {/if}
            </span>
            <span><b>Print Processor:</b>{printer.print_processor}</span>
            <span><b>Notes: </b>{printer.notes}</span>
        </div>
    </section>
{/if}

<style>

    section {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        padding: 1rem;
    }

    header {
        display: flex;
        justify-content: center;
        position: relative;
        gap: 5px;
    }

    div {
        display: flex;
        flex-direction: column;
        gap: 5px;
    }

    span {
        display: flex;
        gap: 10px;
    }

    a {
        color: var(--color-text)
    }

    a:visited {
        color: var(--color-text);
    }

    a:hover {
        color: var(--color-text-opacity-80)
    }


</style>


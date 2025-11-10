<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import { copyToClip, type CopyState } from '$lib/helper/copy.svelte';
	import Printer from '@components/Cards/Printer.svelte';
	import CopyButton from '@components/CopyButton.svelte';
	import PageLoading from '@components/PageLoading.svelte';
	import PrinterStatus from '@components/Printer/PrinterStatus.svelte';
	import type { PrinterSimpleInfo } from '@t/printer';
	import { onMount } from 'svelte';

    let printer: PrinterSimpleInfo | null = $state(null)
    let isOnline: boolean = $state(false)
    let eRecordPrinter: boolean = $state(false)
    let relatedPrinters: PrinterSimpleInfo[] = $state([])
    let loading: boolean = $state(true)

    let {
        data
    } : {
        data: {name: string}
    } = $props()

    onMount(async () => {
        await PageInformation()

    })


    let copyState: CopyState = $state({
        copied: "",
        timeout: null
    })

    $effect(() => {
        void (async () => {
            console.log(data.name)

            await PageInformation()
        })();
    })


    const PageInformation = async () => {
        loading = true
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
        loading = false
        res = await fetch(`http://localhost:8000/api/printer/ping/${printer.ip.replace(/_.*/, "")}`)




        const str = await res.text();
        isOnline = str.includes("Successfully")
        eRecordPrinter = printer.server.includes("eRcd-CPA")

        res = await fetch(`http://localhost:8000/api/printer/related/${printer.ip.replace(/_.*/, "")}`)

        relatedPrinters = await res.json() as PrinterSimpleInfo[]

    }


</script>




{#if printer != null && !loading}
    <section>
        <header>
            <h1>{`\\\\${printer.server}\\${printer.queue}`}</h1>
            <sup><PrinterStatus {isOnline} {eRecordPrinter} ip={printer.ip}/></sup>
        </header>
        <div class="main-printer">
            <CopyButton label={"MODEL"} value={printer.model ? printer.model  : "NA"}/>
            {#if printer.ip}
                {#if isOnline && !eRecordPrinter}
                    <span><button onclick={() => copyToClip(printer ? printer.ip : "NA", copyState)}><b>IP: </b></button><a href={`http://${printer.ip.replace(/_.*/, "")}`}>{printer.ip != copyState.copied ? printer.ip : "Copied"}</a></span>
                {:else}
                    <CopyButton label={"IP"} value={printer.ip}/>
                {/if}
            {/if}
                        <CopyButton label={"LOCATION"} value={printer.location ? printer.location : "NA"} />
            <CopyButton label={"PRINT_PROCESSOR"} value={printer.print_processor ? printer.print_processor : "NA"}/>
            <CopyButton label={"NOTES"} value={printer.notes ? printer.notes : "NA"}/>
        </div>


        {#if relatedPrinters.length > 1}   
            <div class="related-printers">
                <h1>Related Printers</h1>
                {#each relatedPrinters as p, idx}
                    {#if printer.queue != p.queue}
                        <Printer printer={p} {idx}/>
                    {/if}
                {/each}
            </div>
        {/if}
    </section>

{:else} 
<PageLoading/>
{/if}

<style>

    button {
        background: none;
        border: none;
        color: var(--color-text);
        font-size: 15px;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    span {
        display: flex;
        align-items: center;
        gap: 10px;
    }

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

    div.main-printer {
        display: flex;
        flex-direction: column;
        gap: 10px;
        padding: 2rem;
        border-radius: 10px;
        min-width: 50%;
        background: var(--color-surface);
    }

    div.related-printers {
        display: flex;
        flex-direction: column;
        width: 75%;
        gap: 1rem;
        margin-top: 5rem;
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


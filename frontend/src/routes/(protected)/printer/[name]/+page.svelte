<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
    import { Copy } from '$lib/types/copy';
	import { isLoggedIn } from '$lib/stores/login';
	import PrinterCard from '$lib/components/cards/Printer.svelte';
	import CopyButton from '$lib/components/CopyButton.svelte';
	import PageLoading from '$lib/components/loading/PageLoading.svelte';
	import PrinterStatus from '$lib/components/printer/PrinterStatus.svelte';
	import {Printer} from '$lib/types/printer';
	import { onMount } from 'svelte';

    let printer: Printer.CardInfo | null = $state(null)
    let isOnline: boolean = $state(false)
    let eRecordPrinter: boolean = $state(false)
    let relatedPrinters: Printer.CardInfo[] = $state([])
    let loading: boolean = $state(true)

    let {
        data
    } : {
        data: {name: string}
    } = $props()

    onMount(async () => {
        await PageInformation()

    })


    let copyState: Copy.State = $state({
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
        

        printer = await res.json() as Printer.CardInfo
        loading = false
        res = await fetch(`http://localhost:8000/api/printer/ping/${printer.ip.replace(/_.*/, "")}`)

        isOnline = res.status == 200
        eRecordPrinter = printer.server.includes("eRcd-CPA")

        res = await fetch(`http://localhost:8000/api/printer/related/${printer.ip.replace(/_.*/, "")}`)

        relatedPrinters = await res.json() as Printer.CardInfo[]

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
                    <span><button onclick={() => Copy.ToClipboard(printer ? printer.ip : "NA", copyState)}><b>IP: </b></button><a href={`http://${printer.ip.replace(/_.*/, "")}`}>{printer.ip != copyState.copied ? printer.ip : "Copied"}</a></span>
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
                        <PrinterCard item={p} {idx}/>
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


    @media (max-width: 800px) {
        header {
            word-break: break-all;
            flex-direction: column;
        }

        div.related-printers {
            width: 90%;
        }
    }


</style>


<script lang="ts">
	import { Icons } from "$lib/managers/icons";
    import {Printer} from "$lib/types/printer";
	import CopyAllButton from "../buttons/CopyAllButton.svelte";
	import CopyButton from "../buttons/CopyButton.svelte";
	import Card from "./Card.svelte";
	import { PrinterStateClass } from "./states/PrinterState.svelte";


    let {
        item,
        idx
    } : {
        item: Printer.CardInfo
        idx: number
    } = $props()

    let PrinterState: PrinterStateClass = new PrinterStateClass(item)

</script>

<Card {idx} >
    {#snippet header()}
        {@render Link()}
        <CopyAllButton icon={Icons.COPY} copiedIcon={Icons.COPY_SUCCESSFUL} copyTemplate={PrinterState.copyTemplate} />
    {/snippet}
    {#snippet body()}
        {@render Content()}
    {/snippet}
</Card>


{#snippet Link()}
    <a href={PrinterState.pageLink}  data-title={"Go to Printer Page"}> 
        <img src={Icons.PRINTER} alt="Printer Icon Needed">
    </a>
{/snippet}


<!-- * Renders the object information in copy format -->
{#snippet Content()}
    <div id="content">
        <CopyButton value={PrinterState.fullName} category={"title"}/>
        <div id="attributes">
            {#if PrinterState.model} 
                <CopyButton value={PrinterState.model} label={"MODEL"}/>
            {/if}
            {#if PrinterState.ip} 
                <CopyButton value={PrinterState.ip} label={"IP"}/>
            {/if}
            {#if PrinterState.print_processor} 
                <CopyButton value={PrinterState.print_processor} label={"PROCESSOR"}/>
            {/if}
            {#if PrinterState.location} 
                <CopyButton value={PrinterState.location} label={"LOCATION"}/>
            {/if}
            {#if PrinterState.notes} 
                <CopyButton value={PrinterState.readableNotes} label={"NOTES"}/>
            {/if}
        </div>
    </div>
    
{/snippet}



<style >

    div#attributes {
        padding-left: 10px;
        margin-top: 4px;
    }

    img {
        width: 20px;
    }

    a {
        color: var(--color-text);
    }

    a:hover::after:visited {
        color: var(--color-text);
    }

    a:hover::after {
        content: attr(data-title);
        position: absolute;
        top: -20px;
        left: 20px;
        background-color: #333;
        color: var(--color-text);
        padding: 6px 10px;
        border-radius: 4px;
        font-size: 12px;
        white-space: nowrap;
        box-shadow: 0 4px 6px rgba(0,0,0,0.1);
        z-index: 100;
    }

    
    a {
        opacity: 0.8;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    div#content {
        padding-left: 1.5rem;
        padding-top: 0.5rem;
        padding-bottom: 1.5rem;
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



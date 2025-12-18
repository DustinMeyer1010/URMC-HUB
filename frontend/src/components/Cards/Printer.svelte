<script lang="ts">
    import type { PrinterSimpleInfo } from "@t/printer";
    import Icon from '$lib/assets/double-left-arrow-primary.png';
	import CopyAllButton from "@components/CopyAllButton.svelte";
	import CopyButton from "@components/CopyButton.svelte";


    let {
        item,
        idx
    } : {
        item: PrinterSimpleInfo
        idx: number
    } = $props()

    const printerName = `${item.server}?queue=${item.queue}`;
    const copyText: string = `Name: \\\\${item.server}\\${item.queue}\nModel: ${item.model}\nIP: ${item.ip}\nPrint_Processor: ${item.print_processor}\nLocation: ${item.location}\nNotes: ${item.notes}`;
    const printQueueName: string = `\\\\${item.server}\\${item.queue}`;

</script>


<!-- * Renders main content -->
<ul style="--delay: {Math.min(idx * 50, 2000)}ms">
    <CopyAllButton {copyText} />
    <a href={`/printer/${printerName}`}> <img src={Icon} alt=""></a>
    <CopyButton value={printQueueName} fontSize={15} marginBottom={15}/>
    {@render ObjectContent()}
</ul>


<!-- * Renders the object information in copy format -->
<!-- NOTE: Slice starts at two to skip server & queue -->
{#snippet ObjectContent()}
    {#each Object.entries(item).slice(2) as key}
        {#if key[1]}
            <CopyButton value={key[1]} label={key[0]} />
        {/if}
    {/each}
{/snippet}



<style >

    img {
        width: 25px;
    }

    a {
        opacity: 0.8;
        position: absolute;
        display: flex;
        justify-content: center;
        align-items: center;
        top: 0.3rem;
        right: -0.5rem;
        width: 50px;
    }

    a img {
        transition: var(--transition-fast);
        transform: rotate(130deg);
    }

    a:hover img {
        transform: rotate(130deg) translate(-3px, -1px);
    }

    ul {
        display: flex;
        flex-direction: column;
        word-break: break-all;
        position: relative;
        gap: 0.5rem;
        border-radius: 10px;
        padding: 1.5rem;
        padding-left: 1.5rem;
        padding-right: 3rem;
        box-sizing: border-box;
        background: var(--color-surface);
        color: var(--color-text);
        opacity: 0;
        transform: translateY(20px);
        animation: slideIn 0.2s ease-out forwards;
        animation-delay: var(--delay);
        margin: 0;
        list-style: none;
    }

    @media (max-width: 400px) {

        ul {
            padding-right: 1rem;

        }
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



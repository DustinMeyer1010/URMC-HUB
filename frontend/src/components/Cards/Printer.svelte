<script lang="ts">
    import type { PrinterSimpleInfo } from "@t/printer";
    import { copyToClip, type CopyState  } from '$lib/helper/copy.svelte'
    import copyAllIcon from '$lib/assets/copy-color-text.png'
    import outIcon from '$lib/assets/double-left-arrow-primary.png';


    let copyState: CopyState = $state({
        copied: "",
        timeout: null
    })
    let {
        item,
        idx
    } : {
        item: PrinterSimpleInfo
        idx: number
    } = $props()

    const printerName = `${item.server}?queue=${item.queue}`;

    const allCopyText: string = `Name: \\\\${item.server}\\${item.queue}\nModel: ${item.model}\nIP: ${item.ip}\nPrint_Processor: ${item.print_processor}\nLocation: ${item.location}\nNotes: ${item.notes}`;


</script>

<ul style="--delay: {Math.min(idx * 50, 2000)}ms">
    {#if copyState.copied != allCopyText}
        <button class="copy-all" title="Copy All" onclick={() => copyToClip(allCopyText, copyState)}><img src={copyAllIcon} alt="Copy All"></button>
    {:else}
        <span class="copied-all">ALL COPIED</span>
    {/if}
    <a href={`/printer/${printerName}`}> <img src={outIcon} alt=""></a>
    <li class="name">
        <button
        type="button"
        onclick={() => copyToClip(`\\\\${item.server}\\${item.queue}`, copyState)}>

            {copyState.copied == `\\\\${item.server}\\${item.queue}` ? "Copied" : `\\\\${item.server}\\${item.queue}`}
        </button>
    </li>
    {#each Object.entries(item).slice(2) as key}
        {#if key[1]}
            <li> 
                <button
                type="button"
                onclick={() => copyToClip(key[1], copyState)}>
                        <b>{key[0].toUpperCase()}:</b>
                        {copyState.copied == key[1] ? "Copied" : key[1]}
                </button>
            </li>
        {/if}
    {/each}
</ul>



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

    
    button {
        background: none;
        border: none;
        padding: 0;
        font: inherit;
        color: inherit;
        cursor: pointer;
    }

    button:focus,
    button:active{
        outline: none;
    }

    span.copied-all {
        font-size: 10px;
        position: absolute;
        top: 0.4rem;
        left: 0.4rem;
        opacity: 0.3;
        z-index: -1;
        color: var(--color-text)
    }

    button.copy-all {
        position: absolute;
        top: 0.2rem;
        left: 0.2rem;
        opacity: 0.3;
        z-index: -1;
    }

    button.copy-all img {
        width: 20px;
    }

    button.copy-all:hover img {
        transform: scale(1.1);
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


    li.name {
        font-weight: bold;
        font-size: 18px;
        margin-bottom: 1rem;
    }

    li {
        text-align: left;
    }

    button {
        background: none;
        border: none;
        padding: 0;
        font: inherit;
        color: inherit;
        cursor: pointer;
    }

    button:active,
    button:focus { 
        border:none;
        outline: none;
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



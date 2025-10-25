<script lang="ts">
    import type { PrinterSimpleInfo } from "@t/printer";
    import { copyToClip, type CopyState  } from '$lib/helper/copy.svelte'


    let copyState: CopyState = $state({
        copied: "",
        timeout: null
    })

    let {
        printer,
        idx
    } : {
        printer: PrinterSimpleInfo
        idx: number
    } = $props()


</script>


    <ul style="--delay: {Math.min(idx * 50, 2000)}ms">
        <li class="title">
            <button
            type="button"
            onclick={() => copyToClip(`\\\\${printer.server}\\${printer.queue}`, copyState)}>
                \\{printer.server}\{printer.queue} 
                {#if copyState.copied == `\\\\${printer.server}\\${printer.queue}`}
                    Copied
                {/if}
            </button>
        </li>
        {#each Object.entries(printer).slice(2) as key}
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



<style>

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

    ul {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
        border-radius: 10px;
        padding: 1.5rem;
        box-sizing: border-box;
        background: var(--background-surface);
        color: var(--text);
        opacity: 0;
        transform: translateY(20px);
        animation: slideIn 0.2s ease-out forwards;
        animation-delay: var(--delay);
        margin: 0;
        list-style: none;
    }

    li.title {
        font-weight: bold;
        font-size: 18px;
        margin-bottom: 1rem;
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



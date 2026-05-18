<script lang="ts">
	import { Copy } from "$lib/types/copy";
	import { fade } from "svelte/transition";

    let copyState: Copy.State = $state(Copy.EMPTY_COPY_STATE)

    let {
        icon,
        copiedIcon = icon,
        copyTemplate,
        dataTitle = "Copy All"
    } : {
        icon: string,
        copiedIcon?: string
        copyTemplate: string
        dataTitle?: string
    } = $props()

</script>

<button 
    class:copied={copyState.copied != copyTemplate} 
    class="copy-all" 
    title="Copy All" 
    data-title={dataTitle}
    onclick={() => Copy.ToClipboard(copyTemplate, copyState)}
>
        {#key copyState.copied}
            <img 
                in:fade={{ duration: 50}}
                out:fade={{duration: 50}}
                src={copyState.copied != copyTemplate ? icon : copiedIcon} 
                alt="Copy All"
            >
        {/key}
</button>
<style>


    button {
        background: none;
        position: relative;
        border: none;
        padding: 0;
        margin: 0;
        font: inherit;
        color: var(--color-text);
        cursor: pointer;
    }

    button:focus,
    button:active{
        outline: none;
    }

    button.copied {
        color: var(--color-success);
    }

    button.copy-all {
        display: inline-grid;
        place-items: center;
    }

    button.copy-all img {
        grid-area: 1/1;
        height: 30px;
    }

    button:hover::after {
        content: attr(data-title);
        position: absolute;
        top: -30px;
        left: 15px;
        background-color: #333;
        color: var(--color-text);
        padding: 6px 10px;
        border-radius: 5px;
        font-size: 12px;
        white-space: nowrap;
        box-shadow: 0 4px 6px rgba(0,0,0,0.1);
        z-index: 100;
        opacity: 0.8;
    }

    button.copy-all:hover img {
        transform: scale(1.1);
    }


    button:active,
    button:focus { 
        border:none;
        outline: none;
    }
</style>


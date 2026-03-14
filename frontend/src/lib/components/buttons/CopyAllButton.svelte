<script lang="ts">
	import { Copy } from "$lib/types/copy";

	import { fade } from "svelte/transition";

    let copyState: Copy.State = $state(Copy.EMPTY_COPY_STATE)

    let {
        icon,
        copiedIcon = icon,
        copyTemplate
    } : {
        icon: string,
        copiedIcon?: string
        copyTemplate: string
    } = $props()

</script>

<button 
    class:copied={copyState.copied != copyTemplate} 
    class="copy-all" 
    title="Copy All" 
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
        border: none;
        padding: 0;
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
        top: 0.2rem;
        left: 0.2rem;
    }

    button.copy-all img {
        grid-area: 1/1;
        width: 25px;
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


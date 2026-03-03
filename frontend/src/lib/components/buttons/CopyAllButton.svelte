<script lang="ts">
	import { Copy } from "$lib/types/copy";
    import Icon from "$lib/assets/copy-color-text.png"
    let { copyTemplate } : {copyTemplate: string} = $props()

    let copyState: Copy.State = $state(Copy.EMPTY_COPY_STATE)

</script>

{#if copyState.copied != copyTemplate}
    <button class="copy-all" title="Copy All" onclick={() => Copy.ToClipboard(copyTemplate, copyState)}><img src={Icon} alt="Copy All"></button>
{:else}
    <span class="copied-all">ALL COPIED</span>
{/if}


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

    span.copied-all {
        font-size: 15px;
        position: absolute;
        top: 0.4rem;
        left: 0.4rem;
        opacity: 0.3;
        color: var(--color-text)
    }

    button.copy-all {
        position: absolute;
        top: 0.2rem;
        left: 0.2rem;
        opacity: 0.3;
    }

    button.copy-all img {
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


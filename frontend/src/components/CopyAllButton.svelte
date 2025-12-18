<script lang="ts">
	import { copyToClip, type CopyState } from "$lib/helper/copy.svelte";
    import Icon from "$lib/assets/copy-color-text.png"
    let { copyText } : {copyText: string} = $props()

let copyState: CopyState = $state({
        copied: "",
        timeout: null
    })

</script>

{#if copyState.copied != copyText}
    <button class="copy-all" title="Copy All" onclick={() => copyToClip(copyText, copyState)}><img src={Icon} alt="Copy All"></button>
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
</style>
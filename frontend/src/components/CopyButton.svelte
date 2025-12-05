<script lang="ts">
	import { copyToClip, type CopyState } from "$lib/helper/copy.svelte";

 

    let copyState: CopyState = $state({
        copied: "",
        timeout: null
    })

    let {
        value,
        label = "",
    } : {
        value: string
        label: string
    } = $props()
    

</script>



<button onclick={() => copyToClip(value, copyState)}>
    {#if label != ""}
        <b>{label}:</b>
    {/if}
    <span class:header={label == ""}>{@html value == copyState.copied ? "Copied" : value}</span>
</button>




<style>
    b {
        font-size: 15px;
    }

    button {
        margin: 0;
        background: transparent;
        border: none; 
        padding: 0;
        color: var(--color-text);
        text-align: left;
        font-size: 15px;
        font-family: Roboto;
        word-break: break-all;
        width: 100%;
    }

    span.header {
        font-size: 20px;
        font-weight: bold;
    }

    @media (max-width: 500px) {
        button {
            display: flex;
            gap: 5px;
            margin-bottom: 5px;
            flex-direction: column;
        }
    }
</style>
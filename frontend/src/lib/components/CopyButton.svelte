<script lang="ts">
    import { Copy } from "$lib/types/copy"

 

    let copyState: Copy.State = $state(Copy.EMPTY_COPY_STATE)

    let {
        value,
        label = "",
        fontSize = 15,
        marginBottom = 0
    } : {
        value: string
        label?: string
        fontSize?: number
        marginBottom?: number
    } = $props()

    label = label.toUpperCase()
    

</script>



<button style={`--margin-bottom: ${marginBottom}px`} onclick={() => Copy.ToClipboard(value, copyState)}>
    {#if label != ""}
    <b style="--font-size: {fontSize}px">{label}:</b>
    {/if}
    <span style="--font-size: {fontSize}px" class:header={label == ""}>{value == copyState.copied ? "Copied" : value}</span>
</button>




<style>
    b {
        font-size: var(--font-size);
    }

    button {
        margin: 0;
        margin-bottom: var(--margin-bottom);
        background: transparent;
        border: none; 
        padding: 0;
        color: var(--color-text);
        text-align: left;
        font-size: var(--font-size);
        font-family: Roboto;
        word-break: break-all;
        width: 100%;
    }

    span.header {
        font-size: var(--font-size);
        font-weight: bold;
    }

    @media (max-width: 350px) {
        button {
            display: flex;
            gap: 5px;
            margin-bottom: 5px;
            flex-direction: column;
        }
    }
</style>
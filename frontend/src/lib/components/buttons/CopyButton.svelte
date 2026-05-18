<script lang="ts">
    import { Copy } from "$lib/types/copy"

 

    let copyState: Copy.State = $state(Copy.EMPTY_COPY_STATE)
    type kind = "title" | "list-item" | "drive-group"

    let {
        value,
        label = "",
        style = "",
        category = "list-item",
    } : {
        value: string
        style?: string
        label?: string
        category?: kind
    } = $props()

    label = label.toUpperCase()

    let updatedValue: string = $state(value)
    

</script>

<button {style} id={category} onclick={() => Copy.ToClipboard(updatedValue, copyState)} title="Click to Copy">
    {#if category == "title" || category == "drive-group"}
        <b id={category}>
            <span>{updatedValue == copyState.copied ? "COPIED" : updatedValue}</span>
        </b>
    {:else}
        <span>
            <b id={category}>
                {label}:
            </b>
            <span>{updatedValue == copyState.copied ? "COPIED" : updatedValue}</span>
        </span>
    {/if}
</button>




<style>


    b#drive-group {
        margin-bottom: 0.5rem;
    }

    button#title{
        font-size: clamp(8px, 2vw + 1rem, 20px);
    }


    button:hover {
        color: var(--color-text)
    }

    button {
        font-size: 20px;
        margin: 0;
        padding: 0.2rem;
        text-align: left;
        background: transparent;
        position: relative;
        border: none; 
        color: var(--color-text);
        font-family: Roboto;
        word-break: keep-all;
        width: 100%;
        transition: 0.3s;
        font-size: clamp(15px, 2vw + 1rem, 18px);
    }

    

</style>
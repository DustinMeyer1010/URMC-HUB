<script lang="ts">
    import { Copy } from "$lib/types/copy"

 

    let copyState: Copy.State = $state(Copy.EMPTY_COPY_STATE)
    type kind = "title" | "list-item"

    let {
        value,
        label = "",
        category = "list-item",
    } : {
        value: string
        label?: string
        category?: kind
    } = $props()

    label = label.toUpperCase()
    

</script>

<button id={category} onclick={() => Copy.ToClipboard(value, copyState)}>
    {#if category == "title"}
        <b id={category}>{value == copyState.copied ? "Copied" : value}</b>
    {:else}
        <span>
            <b id={category}>{label}:</b>
            <span>{value == copyState.copied ? "Copied" : value}</span>
        </span>
    {/if}
</button>




<style>

    button#title {
        margin-bottom: 0.5rem;
        font-size: 18px;
    }

    button {
        margin: 0;
        padding: 0.2rem;
        margin-bottom: var(--margin-bottom);
        background: transparent;
        border: none; 
        color: var(--color-text);
        text-align: left;
        font-size: 15px;
        font-family: Roboto;
        word-break: keep-all;
        width: 100%;
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
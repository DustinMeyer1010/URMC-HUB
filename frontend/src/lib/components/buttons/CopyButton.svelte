<script lang="ts">
    import { Copy } from "$lib/types/copy"

 

    let copyState: Copy.State = $state(Copy.EMPTY_COPY_STATE)
    type kind = "title" | "list-item" | "drive-group"

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

    let updatedValue: string = $state(value)
    

</script>

<button id={category} onclick={() => Copy.ToClipboard(updatedValue, copyState)} title="Click to Copy">
    {#if category == "title" || category == "drive-group"}
        <b id={category}>
            {updatedValue}
            <sub class="copied">{updatedValue == copyState.copied ? "(COPIED)" : ""}</sub>
        </b>
    {:else}
        <span>
            <b id={category}>
                {label}:
            </b>
            <span>{updatedValue}</span>
            <sub class="copied">{updatedValue == copyState.copied ? "(COPIED)" : ""}</sub>
        </span>
    {/if}
</button>




<style>

    button#title {
        margin-bottom: 0.5rem;
        font-size: 18px;
    }


    b#drive-group {
        margin-bottom: 0.5rem;
        font-size: 15px;
    }

    sub.copied {
        font-size: 10px;
    }

    button:hover {
        color: var(--color-primary-hover)
    }

    button {
        font-size: 15px;
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
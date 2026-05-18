<script lang="ts">



	import { onMount } from 'svelte';

    let {
        search,
        loading = $bindable(),
        searchValue = $bindable(),
    } : {
        search: () => void;
        loading: boolean;
        searchValue: string;
    } = $props();

    let disabled = $derived(searchValue.includes("/"))

    onMount(async () => {
        search()
        loading = false

    })

    async function onsubmit(e: SubmitEvent) {
        if (disabled) {
            return
        }

        e.preventDefault()
        search()
    }

</script>

<form {onsubmit}>
    <input oncontextmenu={(e: Event) => {e.preventDefault();searchValue=""}} type="text" bind:value={searchValue}>
    <button disabled={disabled} type="submit">Search</button>
</form>

<style>
    form {
        display: flex;
        flex-wrap: nowrap;
        gap: 1rem;
        width: 100%;

    }

    input {
        border: none;
        backdrop-filter: blur(10px);
        text-align: center;
        background: var(--color-bg-opacity-80);
        border: 2px solid var(--color-surface-lighter);
        padding: 0.5rem;
        border-radius: 5px;
        font-size: 18px;
        width: 500px;
        color: var(--color-text);
        transition: 0.3s ease;
        flex-basis: 80%;
    }

    input:focus {
        outline: none;
    }

    button {
        background: var(--color-bg-opacity-80);
        backdrop-filter: blur(5px);
        border: none;
                box-shadow: 
        inset 5px 10px 10px 2px rgba(255,255,255,0.15),
        inset -5px -5px 10px 2px rgba(0,0,0,0.3),
        5px 5px 10px 2px rgba(0,0,0,0.1);
        border-radius: 20px;
        padding: 0 2rem;
        font-weight: bold;
        color: var(--color-text);
        transition: 0.3s ease;
        flex-basis: 20%;
    }

    button:hover {
        box-shadow: 
        inset -5px -5px 20px 2px rgba(255,255,255,0.1),
        inset 5px 5px 10px 2px rgba(0,0,0,0.5);
    }



</style>
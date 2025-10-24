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

    onMount(async () => {
    await search()
    loading = false

    })

    async function onsubmit(e: SubmitEvent) {
        e.preventDefault()
        await search()
    }

</script>

<form {onsubmit}>
    <input type="text" bind:value={searchValue}>
    <button type="submit">Search</button>
</form>

<style>
    form {
        display: flex;
        flex-wrap: nowrap;
        gap: 1rem;

    }

    input {
        border: none;
        backdrop-filter: blur(10px);
        text-align: center;
        background: rgba(33,35,53, 0.8);
        border: 2px solid var(--color-border);
        padding: 0.5rem;
        border-radius: 5px;
        font-size: 18px;
        width: 300px;
        color: var(--text);
        transition: 0.3s ease;
    }

    input:focus {
        outline: none;
    }

    button {
        background: transparent;
        backdrop-filter: blur(5px);
        border: 2px solid var(--color-primary);
        border-radius: 5px;
        padding: 0 1rem;
        font-weight: bold;
        color: var(--color-primary);
        transition: 0.3s ease;
    }

    button:hover {
        background: var(--color-primary-hover-opacity-20);
    }

    @media (max-width: 601px) {

        input {
            width: 200px;
        }
    }

    @media (max-width: 400px) {

        input {
            width: 100px;
        }

        button {
            padding: 0 0.2rem;
        }
    }
</style>
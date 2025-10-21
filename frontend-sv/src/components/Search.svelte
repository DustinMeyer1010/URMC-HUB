<script lang="ts">


    import { page } from '$app/state';
	import { onMount } from 'svelte';
    

    let searchParams = page.url.searchParams;
    let query = searchParams.get('search') ?? ""
    let searchValue: string = $state(query)


    let {
        search,
        loading = $bindable(),
    } : {
        search: (searchValue: string) => void;
        loading: boolean;
    } = $props();

    onMount(async () => {
    if (searchValue !== "") {
        await search(searchValue)
    }
    loading = false

    })

    function onsubmit(e: SubmitEvent) {
        e.preventDefault()
        search(searchValue)
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
        background: var(--background-surface);
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
        border: 2px solid var(--primary);
        border-radius: 5px;
        padding: 0 1rem;
        font-weight: bold;
        color: var(--primary);
        transition: 0.3s ease;
    }

    button:hover {
        background: var(--primary);
        color: var(--text);
    }

    @media (max-width: 601px) {
        input {
            width: 200px;
        }
    }
</style>
<script lang="ts">


    import { Search } from '$lib/types/search';
    import { fly } from 'svelte/transition'
    import { get } from 'svelte/store';
    import { theme } from '$lib/stores/theme';

    let isFiltersOpen: boolean = $state(true)
    let currentTheme = get(theme)

    let {
        currentFilter,
        switchFilter
    } : {
        currentFilter: Search.Filter,
        switchFilter: (newFilter: Search.Filter) => void
    } = $props();


</script>
    <div>
        {#each Search.ValidFilters as filter, idx}
            {#if isFiltersOpen}
                <button 
                    class="filter-options"
                    style="order {idx}"
                    in:fly={{delay: 80*idx, y:20, x:-20 * idx}}
                    out:fly={{delay: 0, y:20, x:-20 * idx}}
                    class:active={currentFilter == filter}
                    onclick={() => switchFilter(filter as Search.Filter)}>
                    {filter}
                </button>
            {/if}
        {/each}
    </div>  

<style>


    div {
        display: flex;
        justify-content: center;
        align-items: center;
        width: 100%;
        padding: 10px;
        border-radius: 20px;
        gap: 1rem;
        grid-column: span 3;
    }


    button.filter-options {

        box-shadow: 
        inset -5px -5px 20px 2px rgba(255,255,255,0.1),
        inset 5px 5px 10px 2px rgba(0,0,0,0.5);
        transition: 0.3s ease-in-out;
        padding: 1rem 1rem;
        border: 1px solid rgba(255,255,255,0.3);
        width: 150px;
        border-radius: 20px;
        background: var(--color-surface);
        transform-origin: center;
        color: var(--text);
        transition: 0.3s ease;
        font-weight: bold;
        backdrop-filter: blur(10px);
        transition: 0.5s ease;
    }

    button.active {
        background-color: var(--color-surface);
        box-shadow: 
        inset 5px 5px 10px 2px rgba(255,255,255,0.15),
        inset -5px -5px 10px 2px rgba(0,0,0,0),
        5px 5px 10px 2px rgba(0,0,0,0.3);
        color: var(--color-accent);
        border: 1px solid var(--color-accent-hover-opacity-40);
    }


    button:hover {
        box-shadow: 
        inset 5px 10px 10px 2px rgba(255,255,255,0.15),
        inset -5px -5px 10px 2px rgba(0,0,0,0),
        5px 5px 10px 2px rgba(0,0,0,0.3);
    }

    @media (max-width: 800px) {
        button.filter-options {
            width: 120px;
            font-size: 12px;
        }

        div {
            flex-wrap: wrap-reverse;
        }


    }


</style>
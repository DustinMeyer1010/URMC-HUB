<script lang="ts">


    import { Search } from '@t/search'
    import { fly } from 'svelte/transition'
    import filterIconDark from '$lib/assets/filter-color-accent-dark.png'
    import filterIconLight from '$lib/assets/filter-color-accent-light.png'
    import { get } from 'svelte/store';
    import { theme } from '$lib/theme';

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

    <button class="open-filters" onclick={() => isFiltersOpen = !isFiltersOpen}>
        <img src={currentTheme === "dark"? filterIconDark : filterIconLight} alt="filter">
    </button>

<style>


    div {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-wrap: wrap-reverse;
        gap: 1rem;
        width: 80%;
        flex-basis: 100%;
        order: 0;
    }

    .open-filters {
        min-width: 25px;
        border-radius: 10px;
        padding: 10px;
        box-sizing: border-box;
        bottom: 25px;
        left: 50%;
        display: flex;
        justify-content: center;
        align-items: center;
        border: none;
        background: var(--color-bg-opacity-80);
    }



    img {
        width: 20px;
    }

    button.filter-options {
        padding: 0.5rem 1rem;
        width: 150px;
        border: 2px solid var(--color-accent);
        border-radius: 20px;
        background: var(--color-bg-opacity-80);
        transform-origin: center;
        color: var(--color-accent);
        transition: 0.3s ease;
        font-weight: bold;
        backdrop-filter: blur(10px);
        transition: 0.5s ease;
    }

    button.active {
        background: var(--color-accent-focus);
        color: var(--color-bg)
    }

    button.active:hover {
        background: var(--color-accent-focus);
    }

    button:hover {
        background: var(--color-accent-hover-opacity-20);
    }

    @media (max-width: 950px) {
        button.filter-options {
            width: 120px;
            padding: 5px;
            font-size: 12px;
        }


    }


</style>
<script lang="ts">


    import { Search } from '$lib/types/search'
    import { fly } from 'svelte/transition'
    import filterIconDark from '$lib/assets/filter-color-accent-dark.png'
    import filterIconLight from '$lib/assets/filter-color-accent-light.png'
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

        box-shadow: 
        inset -5px -5px 10px 2px rgba(255,255,255,0.1),
        inset 5px 5px 10px 2px rgba(0,0,0,0.4);
        transition: 0.3s ease-in-out;
        padding: 0.7rem 1rem;
        width: 150px;
        border-radius: 20px;
        border: 0.5px solid var(--color-accent-hover-opacity-40);
        background: var(--color-bg-opacity-30);
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
        inset 5px 5px 10px 2px rgba(255,255,255,0.05),
        inset -5px -5px 10px 2px rgba(0,0,0,0.05),
        5px 5px 10px 2px rgba(0,0,0,0.3);
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
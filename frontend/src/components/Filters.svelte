<script lang="ts">


    import type{ Groups } from '@t/filters'
    import { fly } from 'svelte/transition'
    import filterIcon from '$lib/assets/filter-color-accent.png'

    let isFiltersOpen: boolean = $state(true)

    let {
        currentFilter,
        switchFilter
    } : {
        currentFilter: Groups,
        switchFilter: (newFilter: Groups) => void
    } = $props();

    let allFilters: Groups[] = ['USERS','COMPUTERS', "GROUPS", 'PRINTERS', 'DRIVES']


</script>


    <div>
        {#each allFilters as filter, idx}
            {#if isFiltersOpen}
                <button 
                    style="order {idx}"
                    in:fly={{delay: 80*idx, y:20, x:-20 * idx}}
                    out:fly={{delay: 0, y:20, x:-20 * idx}}
                    class:active={currentFilter == filter}
                    onclick={() => switchFilter(filter)}>
                    {filter}
                </button>

            {/if}
        {/each}
    </div>  

    <button class="open-filters" onclick={() => isFiltersOpen = !isFiltersOpen}>
        <img src={filterIcon} alt="filter">
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
        width: 50px;
        bottom: 25px;
        left: 50%;
        display: flex;
        justify-content: center;
        align-items: center;
    }


    img {
        width: 15px;
    }

    button {
        padding: 0.5rem 1rem;
        width: 150px;
        border: 2px solid var(--primary-accent);
        border-radius: 20px;
        background: transparent;
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
        button {
            width: 120px;
            padding: 5px;
            font-size: 12px;
        }


    }


</style>
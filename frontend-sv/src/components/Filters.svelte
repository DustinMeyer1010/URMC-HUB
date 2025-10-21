<script lang="ts">


    import type Groups from '@types/filters'
    import { fly } from 'svelte/transition'
    import filterIcon from '$lib/assets/filter-icon.png'

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


<section>
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


    <button class="open-filters" onclick={() => isFiltersOpen = !isFiltersOpen}>
        <img src={filterIcon} alt="filter">
    </button>
</section>

<style>
    section {
        display: flex;
        gap: 3rem;
    }

    .open-filters {
        width: 50px;
        position: absolute;
        bottom: 20px;
        left: 50%;
        display: flex;
        justify-content: center;
        align-items: center;
        transform: translateX(calc(-50% - 250px));
    }


    img {
        width: 15px;
    }

    button {
        padding: 0.5rem 1rem;
        width: 150px;
        border: 2px solid var(--complement);
        border-radius: 20px;
        background: transparent;
        transform-origin: center;
        color: var(--complement);
        transition: 0.3s ease;
        font-weight: bold;
        backdrop-filter: blur(10px);
    }

    button.active {
        background: rgba(255, 115, 0, 0.2);
    }

    button:hover {
        background: rgba(255, 115, 0, 0.2);
    }

    @media (max-width: 950px) {
        button {
            width: 120px;
            padding: 5px;
            font-size: 12px;
        }

        section {
            flex-wrap: wrap-reverse;
            justify-content: center;
            align-items: center;
            gap: 1rem;
        }

    }

    @media (max-width: 601px) {
        .open-filters {
            transform: translateX(calc(-50% - 200px));
        }
    }

</style>
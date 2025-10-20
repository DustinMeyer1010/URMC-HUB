<script lang="ts">

    import type Groups from '@types/filters'

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
        <button 
            style="order {idx}" 
            class:active={currentFilter == filter} 
            onclick={() => switchFilter(filter)}>
            {filter}
        </button>
    {/each}
</section>



<style>
    section {
        display: flex;
        gap: 3rem;
    }

    button {
        padding: 0.5rem 1rem;
        width: 150px;
        border: 2px solid var(--complement);
        border-radius: 20px;
        background: transparent;
        color: var(--complement);
        transition: 0.3s ease;
        font-weight: bold;
    }

    button.active {
        background: var(--complement);
        color: var(--text)
    }

    button:hover {
        background: var(--complement);
        color: var(--text)
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

</style>
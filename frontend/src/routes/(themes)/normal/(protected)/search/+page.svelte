<script lang="ts">

    import Search from '$lib/components/search/Search.svelte'
    import Filters from '$lib/components/search/Filters.svelte'
    import Cards from '$lib/components/cards/Cards.svelte';

	import { onMount } from 'svelte';
	import CardLoading from '$lib/components/loading/CardLoading.svelte';
	import { SearchStateClass } from './state.svelte';


    let SearchState: SearchStateClass = new SearchStateClass()


    onMount(async () => {
        SearchState.GetURLParams()
        SearchState.SwitchFilter(SearchState.filter)
        SearchState.Search()
    })

</script>

<main>
    {#if SearchState.loading}
        <CardLoading />
    {:else}
        <Cards data={SearchState.data} filter={SearchState.filter}/>
    {/if}

    <section>
        <Filters currentFilter={SearchState.filter} switchFilter={SearchState.SwitchFilter}/>
        <Search search={SearchState.Search} bind:searchValue={SearchState.searchValue} bind:loading={SearchState.loading}/>
    </section>
</main>


<style>

    main {
        width: 100%;
        padding-bottom: 120px;
        display: flex;
        flex-direction: column;
        gap: 1rem;
    }


    section {
        position: fixed;
        display: grid;
        grid-template-columns: 1fr;
        grid-template-rows: auto auto;
        padding-bottom: 20px;
        width: fit-content;
        place-items: center;
        gap: 1rem; 
        bottom: 0px;
        left: 50%;
        transform: translateX(-50%);
    }
    
    @media (max-width: 800px) {
        section {
            width: 100%;
        }

    }
</style>
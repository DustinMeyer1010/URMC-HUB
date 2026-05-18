<script lang="ts">

    import Filters from '$lib/components/search/Filters.svelte'
    import Cards from '$lib/components/cards/Cards.svelte';

	import { onMount } from 'svelte';
	import CardLoading from '$lib/components/loading/CardLoading.svelte';
	import { SearchStateClass } from './state.svelte';
	import type { Search as S } from '$lib/types/search';
	import Search from '$lib/components/search/Search.svelte';


    let {data} : {data: {searchValue: string, filter: S.Filter}} = $props()

    let SearchState: SearchStateClass = new SearchStateClass(data.filter, data.searchValue)


    onMount(async () => {
        SearchState.SwitchFilter(data.filter)
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
        padding-bottom: 20px;
        width: 50%;
        place-items: center;
        gap: 1rem; 
        bottom: 0px;
        left: 50%;
        transform: translateX(-50%);
    }
    
    @media (max-width: 800px) {
        section {
            width: 80%;
        }

    }
</style>
<script lang="ts">

    import Search from '@components/Search.svelte'
    import Filters from '@components/Filters.svelte'
    import Cards from '@components/Cards.svelte';

	import { onMount } from 'svelte';
    import { page } from '$app/state';
	import { getFilter, getSearchValue } from '../../helper';
	import CardLoading from '@components/CardLoading.svelte';
	import { SearchStateClass } from './state.svelte';


    let SearchState: SearchStateClass = new SearchStateClass()


    onMount(() => {
        let urlParams = page.url.searchParams;
        SearchState.searchValue = getSearchValue(urlParams)
        SearchState.filter = getFilter(urlParams)
        SearchState.SwitchFilter(SearchState.filter)
        SearchState.Search()
    })


    const setPreviousSearch = () => {
        localStorage.setItem("searchValue", SearchState.searchValue)
        localStorage.setItem("filter", SearchState.filter)
    }

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
        display: flex;
        justify-content: center;
        align-items: center;
        gap: 2rem;
        flex-wrap: wrap;
        padding-bottom: 20px;
        width: 100%;
        bottom: 0px;
        left: 50%;
        transform: translateX(-50%);
    }
    
    @media (max-width: 800px) {
        section {
            gap: 10px
        }

    }
</style>
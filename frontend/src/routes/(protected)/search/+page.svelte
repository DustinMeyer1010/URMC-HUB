<script lang="ts">

    import Search from '@components/Search.svelte'
    import Filters from '@components/Filters.svelte'
    import Cards from '@components/Cards.svelte';

	import { onMount } from 'svelte';
	import CardLoading from '@components/Loading-Animations/CardLoading.svelte';
	import { SearchStateClass } from './state.svelte';


    let SearchState: SearchStateClass = new SearchStateClass()


    onMount(() => {
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
        padding-top: 10px;
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
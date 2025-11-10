<script lang="ts">

    import Search from '@components/Search.svelte'
    import Filters from '@components/Filters.svelte'
    import Cards from '@components/Cards.svelte';


    import type { Groups } from '@t/filters'
	import type { Results } from '@t/filters';
	import type { AllResults } from '@t/resutls';

    import { goto } from '$app/navigation'
	import { onMount } from 'svelte';
    import { page } from '$app/state';
	import { getFilter, getSearchValue } from '../../helper';
	import CardLoading from '@components/CardLoading.svelte';
    

    let data: AllResults | null = $state(null)
    let items: Results = $state([])
    let filter: Groups = $state("USERS")
    let loading: boolean = $state(true)
    let searchValue: string = $state("")
    let santizedSearch: string = $derived.by(() => {
        return encodeURIComponent(searchValue)
    })


    onMount(() => {
        let urlParams = page.url.searchParams;
        searchValue = getSearchValue(urlParams)
        $inspect(filter)
        filter = getFilter(urlParams)
        switchFilter(filter)
        search()
    })

    const search = async () => {
        if (searchValue == "") {
            loading = false
            return
        }
        loading = true
        data = null
        let response = await fetch(`http://localhost:8000/api/search/all/${santizedSearch}`,{mode: "cors"});
        data = await response.json();

        switchFilter(filter)
        setURL()
        setPreviousSearch()
        loading = false

    }

    const setPreviousSearch = () => {
        localStorage.setItem("searchValue", searchValue)
        localStorage.setItem("filter", filter)
    }


    const setURL = () => {
        goto(`?search=${santizedSearch}&filter=${filter.toLowerCase()}`, { replaceState: true, keepFocus: true, noScroll: true })
    }

    const switchFilter = (newFilter: Groups) => {
        filter = newFilter;
        window.scrollTo(0,0)
        
        if (data === null) {
            return 
        }

        switch (filter) {
            case "COMPUTERS":
                items = data.computers
                break
            case "PRINTERS":
                items = data.printers
                break
            case "USERS":
                items = data.users 
                break
            case "GROUPS":
                items = data.groups
                break
            case "DRIVES":
                items = data.drives
                break
            default:
                filter = "USERS"
                items = data.users 
        }
        setURL()
    }
</script>

<main>
    {#if loading  && data == null}
        <CardLoading />
    {:else}
        <Cards items={items} {filter}/>
    {/if}

    <section>
        <Filters currentFilter={filter} switchFilter={switchFilter}/>
        <Search search={search} bind:searchValue={searchValue} bind:loading={loading}/>
    </section>
</main>




<style>
    div {
        padding: 1rem;
        display: flex;
        flex-direction: column;
        gap: 1rem;
    }

    span {
        background-color: var(--color-surface);
        width: 100%;
        height: 100px;
        border-radius: 10px;
        animation: 3s flash infinite;
        animation-delay: var(--delay);
    }


    @keyframes flash {
        from {
            background-color: var(--color-bg-flash);
        }
        to {
            background-color: var(--color-surface);
        }
    }

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
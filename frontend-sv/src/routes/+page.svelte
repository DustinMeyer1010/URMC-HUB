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

    
    let urlParams = page.url.searchParams;
    let data: AllResults | null = $state(null)
    let items: Results = $state([])
    let filter: Groups = $state(urlParams.get("filter")?.toUpperCase() as Groups ?? "USERS")
    let loading: boolean = $state(true)
    let searchValue: string = $state(urlParams.get('search') ?? "")
    let santizedSearch: string = $derived.by(() => {
        return encodeURIComponent(searchValue)
    })

    


    onMount(() => {
        if (searchValue == "") {
            loading = false
            return
        }
        switchFilter(filter)
        search()
    })

    async function search() {
        if (searchValue == "") {
            loading = false
            return
        }
        loading = true
        data = null
        let response = await fetch(`http://localhost:8000/search/all/${santizedSearch}`,{mode: "cors"});
        data = await response.json();

        switchFilter(filter)
        setURL()
        loading = false
    }

    $inspect(data)

    function setURL() {
        goto(`/?search=${santizedSearch}&filter=${filter.toLowerCase()}`, { replaceState: true, keepFocus: true, noScroll: true })
    }

    function switchFilter(newFilter: Groups) {
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
        <div>
            {#each Array(10).fill(0) as _,x}
                <span style={`--delay: ${250*x}ms`}></span>
            {/each}
        </div>
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
        background-color: var(--background-surface);
        width: 100%;
        height: 100px;
        border-radius: 10px;
        animation: 3s flash infinite;
        animation-delay: var(--delay);
    }


    @keyframes flash {
        from {
            background-color: #4c4e6d;
        }
        to {
            background-color: var(--background-surface);
        }
    }

    main {
        overflow-x: hidden;
        overflow-y: scroll;
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

    @media (max-width: 601px) {

    }


    
</style>
<script lang="ts">

    import Search from '@components/Search.svelte'
    import Filters from '@components/Filters.svelte'
    import Cards from '@components/Cards.svelte';
	import type { ComputerSimpleInfo } from '@types/computer';
    import type { PrinterSimpleInfo } from '@types/printer'
    import type { UserSimpleInfo } from '@types/user';
    import type { GroupSimpleInfo } from '@types/group'
    import type Groups from '@types/filters'
	import type { Results } from '@types/filters';


    let items: Results = $state([])
    let filter: Groups = $state("USERS")
    let data: any | null = $state(null)


    async function onsubmit(e: SubmitEvent, searchValue: string) {
        e.preventDefault()

        let res = await fetch(`http://localhost:8000/search/all/${searchValue}`);
        data = await res.json();

        console.log(data)

        switchFilter(filter)

    }

    function switchFilter(newFilter: Groups) {
        filter = newFilter;
        
        // console.log(filter)
        // console.log(data)
        if (data === null) {
            return 
        }

        switch (filter) {
            case "COMPUTERS":
                items = data.computers as ComputerSimpleInfo[]
                break
            case "PRINTERS":
                items = data.printers as PrinterSimpleInfo[];
                break
            case "USERS":
                items = data.users as UserSimpleInfo[];
                break
            case "GROUPS":
                items = data.groups as GroupSimpleInfo[];
                break
            default:
                filter = "USERS"
                items = data.users as UserSimpleInfo[];
        }
    }



</script>


<main>
    <Cards items={items} {filter}/>
    <section>
        <Filters currentFilter={filter} switchFilter={switchFilter}/>
        <Search {onsubmit}/>
    </section>
</main>





<style>
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
        flex-direction: column;
        justify-content: center;
        align-items: center;
        gap: 2rem;
        flex-wrap: nowrap;
        padding-bottom: 20px;
        bottom: 0px;
        left: 50%;
        width: 100%;
        background: green;
        transform: translateX(-50%);
    }


    
</style>
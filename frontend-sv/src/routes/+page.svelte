<script>
    import Search from '../components/Search.svelte'
    import Filters from '../components/Filters.svelte'

</script>

<h1>Search Page</h1>



<section>
    <Filters/>

    <Search/>
</section>



<style>
    h1 {
        color: var(--text)
    }

    div {

    }

    section {
        position: fixed;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        gap: 2rem;
        flex-wrap: nowrap;
        bottom: 20px;
        left: 50%;
        transform: translateX(-50%);
    }

    
</style>
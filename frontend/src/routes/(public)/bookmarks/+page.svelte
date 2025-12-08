<script lang="ts">
	import { onMount } from "svelte";

    type Bookmark = {
        name: string,
        url: string,
        image_path: string,
        description: string
    }

    let bookmarks: Bookmark[] = $state([])
    let filter: string = $state("")
    let filteredBookmarks: Bookmark[] = $derived.by(() => {
        if (filter === "") {
            return bookmarks
        }

        return bookmarks.filter((bookmark) => {
            return bookmark.name.toLocaleLowerCase().includes(filter) || bookmark.description.toLocaleLowerCase().includes(filter)
        })


        
    })

    onMount(async () => {
        await fetch("http://localhost:8000/api/generic/bookmarks")
        .then(async (res) => {
            bookmarks = await res.json()
        })
    })  
</script>

<section>
    {#each filteredBookmarks as bookmark}
    <a href={bookmark.url}>
        <h1>{bookmark.name}</h1>
        <span>{bookmark.description}</span>
    </a>
    {/each}

    <input type="text" bind:value={filter}>
</section>


<style>

    a,
    a:visited,
    a:hover {
        text-decoration: none;
        color: var(--text-color);
        text-align: center;
    }

    h1 {
        margin: 0;
        padding: 0;
    }

    a {
        padding: 5px;
        background: var(--color-surface);
        border-radius: 10px;
    }

    section {
        padding: 10px;
        display: grid;
        grid-template-columns: 1fr 1fr 1fr;
        gap: 20px;
        padding-bottom: 150px;
    }

    input {
        position: fixed;
        bottom: 50px;
        left: 50%;
        width: 50%;
        max-width: 500px;
        min-width: 100px;
        text-align: center;
        font-size: 18px;
        border-radius: 5px;
        height: 35px;
        transform: translateX(-50%);
        border: 2px solid #b4b8d9;
        background-color: var(--color-bg-opacity-80);
        color: var(--text-color)
    }

    input:focus {
        outline: none;
    }

</style>
<script lang="ts">
	import { goto } from "$app/navigation";
	import { onMount } from "svelte";

    let { data } : { data: {username: string} } = $props()

    type Bookmark = {
        name: string,
        url: string,
        image_path: string,
        description: string
    }

    let bookmarks: Bookmark[] = $state([])
    let filter: string = $state("")
    let currentBookmarks: string = $state("")
    let agent: string = $state("")
    let agentWithBookmarks: string[] = $state([])

    $effect(() => {
        if (currentBookmarks == "Generic") {
            goto("/bookmarks", {replaceState: true, invalidateAll: true})
            return
        }
        goto(`/bookmarks/${currentBookmarks}`, {replaceState: true, invalidateAll: true})
    }) 


    let filteredBookmarks: Bookmark[] = $derived.by(() => {
        if (filter === "") {
            return bookmarks
        }

        return bookmarks.filter((bookmark) => {
            return bookmark.name.toLocaleLowerCase().includes(filter) || bookmark.description.toLocaleLowerCase().includes(filter)
        })

        
    })

    // This will need to call the api to get the users bookmarks
    onMount( async () => {
        agent = localStorage.getItem("agent") ?? ""
        currentBookmarks = data.username

        await fetch("http://localhost:8000/api/bookmarks/all/agents").then(async (res) => {
            agentWithBookmarks = await res.json()
            agentWithBookmarks = agentWithBookmarks.filter((agentB) => agentB != agent)
        })

        await fetch(`http://localhost:8000/api/bookmarks/${data.username}`).then(async (res) => {
            bookmarks = await res.json()
        })
    })  

</script>



<header>
    <div>
        <input placeholder="Search Link" type="text" bind:value={filter}>
        <select bind:value={currentBookmarks}>
            <option value="Generic">Generic Bookmarks</option>
            <option value={`${agent}`}>My Bookmarks</option>
            {#each agentWithBookmarks as agentB}
                <option value={`${agentB}`}>{agentB}</option>
            {/each}
        </select>
    </div>
{@render AddBookmarks()}
</header>

<section>
    {#each filteredBookmarks as bookmark}
    <a href={bookmark.url}>
        <img src="http://localhost:8000/api/db/image/{bookmark.image_path}" alt="">
        <h1>{bookmark.name}</h1>
        <span>{bookmark.description}</span>
    </a>
    {/each}
</section>


{#snippet AddBookmarks()}
    {#if agent != "" && agent == currentBookmarks}
        <button>
            Add Book Mark
        </button>
    {/if}
{/snippet}

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
        padding-top: 100px;
        display: grid;
        grid-template-columns: 1fr 1fr 1fr;
        gap: 20px;
        padding-bottom: 150px;
    }

    header {
        position: fixed;
        display: flex;
        justify-content: space-between;
        align-items: center;
        width: 100%;
        background-color: var(--color-bg-opacity-80);
        backdrop-filter: blur(10px);
        padding: 20px 20px;
        box-sizing: border-box;
        gap: 10px;
        z-index: 3;
    }

    input {
        width: 50%;
        min-width: 600px;
        font-size: 18px;
        border-radius: 10px;
        height: 35px;
        padding: 5px;
        border: 2px solid #b4b8d9;
        background-color: var(--color-bg-opacity-80);
        color: var(--text-color)
    }

    input:focus {
        outline: none;
    }

    select {
        background: var(--color-surface);
        border: none;
        color: var(--text-color);
        padding: 10px;
        font-size: 15px;
        margin: 0;
    }

</style>
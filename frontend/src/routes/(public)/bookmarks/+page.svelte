<script lang="ts">
	import { goto } from "$app/navigation";
	import { onMount } from "svelte";

    type Bookmark = {
        type: string,
        name: string,
        url: string,
        image_path: string,
        description: string
    }

    let bookmarks: Bookmark[] = $state([])
    let filter: string = $state("")
    let currentBookmarks: string = $state("Generic")
    let agent: string | null = $state("")
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

    $inspect(currentBookmarks)

    onMount(async () => {
            agent = localStorage.getItem("agent") ?? ""
            await fetch("http://localhost:8000/api/generic/bookmarks")
            .then(async (res) => {
                    bookmarks = await res.json()
                })

            await fetch("http://localhost:8000/api/bookmarks/all/agents").then(async (res) => {
                agentWithBookmarks = await res.json()
                agentWithBookmarks = agentWithBookmarks.filter((agentB) => agentB != agent)
            })
        }) 

    const OpenDirectory = async (directory: string) => {
        await fetch("http://localhost:8000/api/directory", {
            mode: "cors",
            method: "POST",
            body: JSON.stringify({path: directory})
        })
    }


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
</header>

<section>
    {#each filteredBookmarks as bookmark}
        {#if bookmark.type == "url"}
            <a href={bookmark.url} target="_blank">
                <img src="http://localhost:8000/api/static/image/{bookmark.image_path}" alt="">
                <h1>{bookmark.name}</h1>
                <span>{bookmark.description}</span>
                
            </a>
        {:else}
        <button onclick={() => OpenDirectory(bookmark.url)}> 
            <img src="http://localhost:8000/api/static/image/{bookmark.image_path}" alt="">
            <h1>{bookmark.name}</h1>
            <span>{bookmark.description}</span>
        </button>

        {/if}
    {/each}


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
        display: flex;
        flex-direction: column;
        justify-content: space-evenly;
        align-items: center;
        padding: 20px;
        box-sizing: border-box;
        background: var(--color-surface);
        border-radius: 10px;
        gap: 20px;
        border: 3px solid transparent
    }

    img {
        width: 400px;
        border-radius: 20px;
        margin-top: -100px;
        transition: var(--transition-normal);
    }

    a:hover,
    a:focus {
        outline: none;
        border: 3px solid var(--color-primary);
    }


    a:focus img {
        transform: translateY(-10px);
    }

    a:hover img {
        transform: translateY(-10px);
    }

    section {
        display: grid;
        box-sizing: border-box;
        grid-template-columns: 1fr 1fr 1fr;
        gap: 100px 20px;
        padding: 200px 50px;
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
        height: 100%;
    }

</style>
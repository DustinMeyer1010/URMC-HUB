<script lang="ts">
	import { goto } from "$app/navigation";
	import Bookmark from "@components/Bookmarks/Bookmark.svelte";
	import { onMount } from "svelte";
    import editIcon from "$lib/assets/edit.png";
	import AddForm from "@components/Bookmarks/AddForm.svelte";

    let { data } : { data: {username: string} } = $props()

    export type BookmarkT = {
        id: string,
        name: string,
        url: string,
        image_path: string,
        description: string
    }

    let bookmarks: BookmarkT[] = $state([])
    let filter: string = $state("")
    let currentBookmarks: string = $state("")
    let agent: string = $state("")
    let agentWithBookmarks: string[] = $state([])
    let owner: boolean = $derived(agent == currentBookmarks)
    let editMode: boolean = $state(false)
    let addFormVisiabled: boolean = $state(false)

    const UpdateLocation = async () => {
        if (currentBookmarks == "Generic") {
            await goto("/bookmarks", {replaceState: true, invalidateAll: true})
            return
        }
        await goto(`/bookmarks/${currentBookmarks}`, {replaceState: true, invalidateAll: true})
        await fetch(`http://localhost:8000/api/bookmarks/${data.username}`).then(async (res) => {
            bookmarks = await res.json()
        })
        editMode = false
        return
    }

    const showAddForm = () => {
        addFormVisiabled = true
    }

    const closeForm = () => {
        addFormVisiabled = false
    }


    let filteredBookmarks: BookmarkT[] = $derived.by(() => {
        if (filter === "") {
            return bookmarks
        }

        const lowerFilter = filter.toLocaleLowerCase()
        return bookmarks.filter((bookmark) => {
            return bookmark?.name?.toLocaleLowerCase().includes(lowerFilter) || bookmark?.description?.toLocaleLowerCase().includes(lowerFilter)
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
        <select onchange={UpdateLocation} bind:value={currentBookmarks}>
            <option value="Generic">Generic Bookmarks</option>
            <option value={`${agent}`}>My Bookmarks</option>
            {#each agentWithBookmarks as agentB}
                <option value={`${agentB}`}>{agentB}</option>
            {/each}
        </select>


    </div>
        {#if owner}
            <button class="edit" onclick={() => editMode = !editMode}><img src={editIcon} alt=""/></button>
        {/if}

</header>

{#if addFormVisiabled} 
    <AddForm username={agent} {closeForm}/>
{/if}

<section>
    {@render AddBookmarks()}
    {#each filteredBookmarks as bookmark}
        <Bookmark username={agent} {bookmark} editMode={editMode} />
    {/each}
</section>


{#snippet AddBookmarks()}
    {#if editMode}
        <button class="add-bookmark" onclick={() => showAddForm()}>
            +
        </button>
    {/if}
{/snippet}



<style>


    button.add-bookmark {
        border: 3px solid var(--color-surface);
        background: none;
        color: var(--text-color);
        border-radius: 10px;
        font-weight: bold;
        font-size: 50px;
    }

    button.add-bookmark:hover {
        background: var(--color-surface);
    }

    button.edit img {
        width: 25px;
    }

    button.edit {
        background: none;
        border: none;
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
        font-size: 15px;
        border-radius: 5px;
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
        padding: 5px;
        font-size: 15px;
    }

</style>
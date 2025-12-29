<script lang="ts">
	import Bookmark from "@components/Bookmarks/Bookmark.svelte";
	import { onMount } from "svelte";
    import editIcon from "$lib/assets/edit.png";
	import AddForm from "@components/Bookmarks/AddForm.svelte";
	import { BookmarkStateClass } from "./state.svlete";

    let { data } : { data: {username: string} } = $props()

    let state: BookmarkStateClass = new BookmarkStateClass(data.username)


    // This will need to call the api to get the users bookmarks
    onMount( async () => {
        state.LoggedInAgent = localStorage.getItem("agent") ?? ""
        state.SelectedAgentBookmarks = data.username
        await state.UpdateAgentBookmarksList()
        await state.UpdateBookmarks()
    })  

</script>



<header>
    <div>
        <input placeholder="Search Link" type="text" bind:value={state.Filter}>
        <select onchange={state.UpdateLocation} bind:value={state.SelectedAgentBookmarks}>
            <option value="Generic">Generic Bookmarks</option>
            <option value={`${state.LoggedInAgent}`}>My Bookmarks</option>
            {#each state.AgentsWithBookmarks as agent}
                <option value={`${agent}`}>{agent}</option>
            {/each}
        </select>


    </div>
        {#if state.Owner}
            <button class="edit" onclick={() => state.EditMode = !state.EditMode}><img src={editIcon} alt=""/></button>
        {/if}

</header>

{#if state.AddFormVisiable} 
    <AddForm username={state.LoggedInAgent} closeForm={state.HideForm} updateBookmarks={state.UpdateBookmarks}/>
{/if}

<section>
    {@render AddBookmarks()}
    {#each state.FilteredBookmarks as bookmark}
        <Bookmark username={state.LoggedInAgent} {bookmark} editMode={state.EditMode}/>
    {/each}
</section>


{#snippet AddBookmarks()}
    {#if state.EditMode}
        <button class="add-bookmark" onclick={() => state.ShowForm()}>
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
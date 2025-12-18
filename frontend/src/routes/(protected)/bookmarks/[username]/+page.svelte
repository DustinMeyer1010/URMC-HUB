<script lang="ts">
	import { goto } from "$app/navigation";
	import { onMount } from "svelte";

    type Bookmark = {
        name: string,
        url: string,
        image_path: string,
        description: string
    }

    let bookmarks: Bookmark[] = $state([])
    let filter: string = $state("")
    let currentBookmarks: string = $state("")

    $effect(() => {
        if (currentBookmarks == "Generic") {
            goto("/bookmarks", {replaceState: true, invalidateAll: true})
            return
        }
        
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
    })  
</script>

<header>
<input placeholder="Search Link" type="text" bind:value={filter}>
<select bind:value={currentBookmarks}>
    <option value="Generic">Generic Bookmarks</option>
    <option value={`${""}`}>My Bookmarks</option>
    <option value={`${"brown"}`}>My Bookmarks</option>
</select>
</header>

<section>
    {#each filteredBookmarks as bookmark}
    <a href={bookmark.url}>
        <h1>{bookmark.name}</h1>
        <span>{bookmark.description}</span>
    </a>
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
        justify-content: center;
        align-items: center;
        width: 100%;
        background-color: var(--color-bg);
        padding: 20px 0px;
        box-sizing: border-box;
        gap: 10px;
    }

    input {
        width: 50%;
        max-width: 500px;
        min-width: 100px;
        text-align: center;
        font-size: 18px;
        border-radius: 5px;
        height: 35px;
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
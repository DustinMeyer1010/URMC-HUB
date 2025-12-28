<script lang="ts">
	import type { BookmarkT } from "../../routes/(protected)/bookmarks/[username]/+page.svelte";


    let {username,bookmark, editMode = false} : {username: string,bookmark: BookmarkT, editMode?: boolean} = $props()

    const removeBookmark = async (id: string) => {
        await fetch(`http://localhost:8000/api/bookmarks/${username}/${id}`, {
            method: "DELETE",
            mode: "cors"
        })
    }

</script>

    <div>
        {@render RemoveBookmark(bookmark.id)}
        <a href={bookmark.url.includes("http") ? bookmark.url : "https://" + bookmark.url} target="_blank"> 
            <img src="http://localhost:8000/api/db/image/{bookmark.image_path}" alt="">
            <h1>{bookmark.name}</h1>
            <span>{bookmark.description}</span>
        </a>
    </div>


    {#snippet RemoveBookmark(id: string)}
        {#if editMode}
            <button class="remove-bookmark" onclick={() => removeBookmark(id)}>
                X
            </button>
        {/if}
    {/snippet}  

<style>

    div {
        position: relative;
    }


    button.remove-bookmark {
        position: absolute;
        top: -10px;
        right: -10px;
        height: 30px;
        width: 30px;
        color: var(--text-color);
        border: none;
        background: var(--color-surface-hover);
        border-radius: 30px;
        z-index: 1;
    }

    button.remove-bookmark:hover {
        border: 1px solid var(--color-surface-hover);
        background: var(--color-bg);
        cursor: pointer;
    }



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
        position: relative;
        flex-direction: column;
        justify-content: space-evenly;
        align-items: center;
        padding: 20px;
        min-height: 350px;
        box-sizing: border-box;
        background: var(--color-surface);
        border-radius: 10px;
        gap: 20px;
        border: 3px solid transparent
    }

    img {
        height: 200px;
        background: var(--color-surface);
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
</style>
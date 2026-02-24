<script lang="ts">

    let {username, closeForm, updateBookmarks} : {username: string, closeForm: () => void, updateBookmarks: () => Promise<void>} = $props()

    let form: HTMLFormElement;

    type Bookmark = {
        name: string,
        url: string,
        description: string,
    }

    let newBookmark: Bookmark = $state({
        name: "",
        url: "",
        description: "",
    })


    
    const onsubmit = async (e: SubmitEvent) => {
        e.preventDefault()
        let formData = new FormData(form) 

        formData.append("bookmark", JSON.stringify(newBookmark))

        await fetch(`http://localhost:8000/api/bookmark/${username}`, {
            method: "POST",
            mode: "cors",
            body: formData
        })

        closeForm()
        await updateBookmarks()
    }

</script>

<div>

    <form bind:this={form} {onsubmit}>
        <button class="close" onclick={closeForm}>X</button>
        <label for="name">Name</label>
        <input bind:value={newBookmark.name} class="text" type="text" id="name" name="name" required>
        <label for="url" >URL</label>
        <input bind:value={newBookmark.url} class="text" type="text" id="url" name="url" required>
        <label for="description">Description</label>
        <input bind:value={newBookmark.description} class="text" type="text" id="description" name="description">
        <label for="photo">Photo</label>
        <input type="file" id="image" name="image" required>
        <button type="submit">Add Bookmark</button>
    </form>`
</div>


<style>

    form {
        position: absolute;
        display: flex;
        flex-direction: column;
        padding: 20px;
        border-radius: 10px;
        background: var(--color-bg);
        left: 50%;
        top: 50%;
        transform: translate(-50%, -50%);
        gap: 10px;
        box-shadow: 0px 10px 10px rgba(0, 0, 0, 0.3);
    }

    input.text {
        width: 300px;
        height: 15px;
        font-size: 15px;
        padding: 8px;
        color: var(--color-text);
        background: var(--color-surface);
        border: 1px solid var(--color-text);
        border-radius: 5px;
    }

    input.text:focus {
        outline: none;
    }

    button {
        background: var(--color-surface);
        border: none;
        color: var(--color-text);
        padding: 8px;
        font-weight: bold;
        border-radius: 5px;
    }

    button.close {
        align-self: flex-end;
    }

    button:hover {
        background: var(--color-surface-hover);
        cursor: pointer;
    }

    div {
        position: absolute;
        top:0;
        left: 0;
        z-index: 3;
        width: 100%;
        height: 100%;
        background: none;
        backdrop-filter: blur(2px);
    }
</style>
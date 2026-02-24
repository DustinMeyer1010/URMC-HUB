<script lang="ts">
	import FileUpload from "$lib/components/files/FileUpload.svelte";
	import { allowedExtensions, BulkLookUpStateClass } from "./state.svelte";
	import GenerationLoading from "$lib/components/loading/GenerationLoading.svelte";

    let BulkLookupState: BulkLookUpStateClass = new BulkLookUpStateClass()


</script>


<section>
    {#if BulkLookupState.loading}
        <GenerationLoading />
    {:else}
        <FileUpload extensions={allowedExtensions} updateFiles={BulkLookupState.updateFiles}/>
        <h1>Select File</h1>
        <h1>OR</h1>
        <h1>Enter Manually</h1>
        <textarea bind:value={BulkLookupState.textValues} oninput={BulkLookupState.parseTextValueInput}></textarea>
        <button onclick={ () => BulkLookupState.lookupUsers()}>Get Users</button>
    {/if}

</section>



<style>



    section {
        display: flex;
        gap: 1rem;
        flex-direction: column;
        align-items: center;
        padding-top: 50px;
        width: 100%;
    }

    textarea {
        background: var(--color-surface);
        border: none; 
        color: var(--color-text);
        padding: 0.3rem;
        width: 90%;
        height: 200px;
        white-space: pre-wrap;
    }

    textarea:focus {
        outline: none;
    }

    button {
        background: var(--color-surface);
        border: 2px solid var(--color-surface);
        padding: 0.5rem 2rem;
        border-radius: 5px;
        color: var(--color-text)
    }

    button:hover {
        background: none;
    }

    h1 {
        padding: 0;
        margin: 0;
        font-size: 15px;
    }

</style>
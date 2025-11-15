<script lang="ts">
	import FileUpload from "@components/FileUpload.svelte";
	import { allowedExtensions, BulkLookUpStateClass } from "./state.svelte";

    let BulkLookupState: BulkLookUpStateClass = new BulkLookUpStateClass()

    $inspect(BulkLookupState)

</script>


<section>

  <FileUpload extensions={allowedExtensions} updateFiles={BulkLookupState.updateFiles}/>
  <h1>OR</h1>
  <span>Manully enter values</span>
    <textarea bind:value={BulkLookupState.textValues} oninput={BulkLookupState.parseTextValueInput}></textarea>
</section>



<style>

    section {
        display: flex;
        flex-direction: column;
        align-items: center;
        padding-top: 50px;
    }

    textarea {
        background: var(--color-surface);
        border: none; 
        color: var(--color-text);
        padding: 0.3rem;
        width: 90%;
        height: 400px;
        white-space: pre-wrap;
    }

    textarea:focus {
        outline: none;
    }
</style>
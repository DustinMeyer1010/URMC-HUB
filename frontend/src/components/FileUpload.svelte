<script lang="ts">
	import { onMount } from "svelte";
	import { FileUploadStateClass } from "./FileUploadState.svelte";

    let {
        extensions,
        updateFiles
    } : {
        extensions: string[]
        updateFiles: (selectedFiles: File[]) => void
    } = $props()

    let FileUploadState: FileUploadStateClass = new FileUploadStateClass(extensions)


    onMount(() => {
        FileUploadState.createInputField()
    })

    $effect(() => {
        updateFiles(FileUploadState.files)
    })


</script>

<section>
    <!-- svelte-ignore a11y_click_events_have_key_events -->
    <div
    role="button"
    tabindex="0"
    onclick={FileUploadState.openFileMenu}
    class="dropzone {FileUploadState.dragging ? 'dragging' : ''}"
    ondragover={FileUploadState.handleDragOver}
    ondragleave={FileUploadState.handleDragLeave}
    ondrop={FileUploadState.handleDrop}
    >
        <strong>Drop files here</strong><br>
        or click to browse
    </div>
    {#if FileUploadState.error}
        {FileUploadState.error}
        <span>Files must be of type .csv, .xlsx, .txt</span>
    {/if}
    {#if FileUploadState.files.length}
        <h3>Uploaded Files</h3>

        <ul>
            {#each FileUploadState.files as file}
            <li>{file.name} — {file.size} bytes</li>
            {/each}
        </ul>
    {/if}
</section>





<style>
    section {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        width: 50%;
        gap: 20px;
    }

  .dropzone {
    width: 100%;
    padding: 100px;
    border: 4px dashed var(--color-surface);
    border-radius: 14px;
    background: transparent;
    text-align: center;
    transition: background 0.2s, border-color 0.2s;
    cursor: pointer;
  }

  .dropzone.dragging {
    border-color: var(--color-text);
    background: var(--color-surface-lighter);
  }
</style>
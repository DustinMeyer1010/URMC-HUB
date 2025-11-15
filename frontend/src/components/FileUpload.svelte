<script lang="ts">
	import { onMount } from "svelte";


    let inputField: HTMLInputElement | undefined = $state(undefined);
    let dragging: boolean = $state(false)
    let error: string = $state("")
    let files: File[] = $state([])

    let {
        extensions,
        updateFiles
    } : {
        extensions: string[]
        updateFiles: (selectedFiles: File[]) => void
    } = $props()

    onMount(() => {
        inputField = document.createElement("input")
        inputField.type = "file"
        inputField.multiple = true;
        inputField.accept = extensions.join(",");

        inputField.addEventListener("change", (e) => {
            const selectedFiles = (e.target as HTMLInputElement).files;
            if (selectedFiles) {
                files = selectedFiles
            }
        })
    })

    $effect(() => {
        updateFiles(files)
    })

    const isAllowed = (file: File) => {
        const ext = file.name.toLowerCase().slice(file.name.lastIndexOf("."))
        return extensions.includes(ext);
    }

    const OpenMenu = () => {
        if (inputField) {
            inputField.click()
        }
    }

    const handleDrop = (e: DragEvent) => {
        e.preventDefault()
        error = ""
        const dropped = Array.from(e.dataTransfer?.files ?? []);
        const valid = dropped.filter(isAllowed)
        const invalid = dropped.filter(f => !isAllowed(f))

        if (invalid.length) {
            error = `Invalid files: ${invalid.map(f => f.name).join(", ")}`;
        }

        files = valid
        dragging = false
    }

    const handleDragOver = (e: DragEvent) => {
        e.preventDefault();
        dragging = true
    }

    const handleDragLeave = () => {
        dragging = false
    }




</script>

<section>
    <!-- svelte-ignore a11y_click_events_have_key_events -->
    <div
    role="button"
    tabindex="0"
    onclick={OpenMenu}
    class="dropzone {dragging ? 'dragging' : ''}"
    ondragover={handleDragOver}
    ondragleave={handleDragLeave}
    ondrop={handleDrop}
    >
        <strong>Drop files here</strong><br>
        or click to browse
    </div>
    {#if error}
        {error}
        <span>Files must be of type .csv, .xlsx, .txt</span>
    {/if}
    {#if files.length}
        <h3>Uploaded Files</h3>

        <ul>
            {#each files as file}
            <li>{file.name} — {file.size} bytes</li>
            {/each}
        </ul>
    {/if}
</section>





<style>
  .dropzone {
    width: 100%;
    max-width: 400px;
    padding: 40px;
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
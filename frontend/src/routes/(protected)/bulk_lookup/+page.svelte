<script lang="ts">
	import FileUpload from "@components/FileUpload.svelte";


    let names: string = $state("")
    let files: File[] = $state([])

    const allowedExtensions = [".txt", ".xlsx", ".csv"];

    $inspect(files)

    const updateFiles = (selectedFiles: File[]) => {
      files = selectedFiles
    }


</script>


<section>

  <FileUpload extensions={allowedExtensions} updateFiles={updateFiles}/>
    <textarea bind:value={names} oninput={() => {
        names = names.split("\n").map(line => line.trimStart().replaceAll("\t", " ")).join("\n")
    }}></textarea>
</section>



<style>



    section {
        display: flex;
        flex-direction: column;
        align-items: center;
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
</style>
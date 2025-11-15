export class FileUploadStateClass {
    
    dragging: boolean = $state(false)
    error: string = $state("")
    files: File[] = $state([]) 
    extensions: string[] = []
    inputField: HTMLInputElement | undefined = $state(undefined);

    constructor(ext: string[]) {
        this.extensions = ext
    }

    createInputField = () => {
this.inputField = Object.assign(document.createElement("input"), {
    type: "file",
    multiple: true,
    accept: this.extensions.join(",")
});

this.inputField.addEventListener("change", (e) => {
    const files = (e.target as HTMLInputElement).files;
    if (files) this.files = Array.from(files);
});
    }


    openFileMenu = () => {
        if (this.inputField) {
            this.inputField.click()
        }
    }



    isAllowed = (file: File): boolean => {
        const ext = file.name.toLowerCase().slice(file.name.lastIndexOf("."))
        return this.extensions.includes(ext);
    }

    handleDrop = (e: DragEvent) => {
        e.preventDefault()
        this.error = ""
        const dropped = Array.from(e.dataTransfer?.files ?? [])
        const valid = dropped.filter((f) => this.isAllowed(f))
        const invalid = dropped.filter((f) => !this.isAllowed(f))

        if (invalid.length) {
            this.error = `Invalid files: ${invalid.map(f => f.name).join(", ")}`;
        }

        this.files = valid
        this.dragging = false
    }

    handleDragOver = (e: DragEvent) => {
        e.preventDefault()
        this.dragging = true
    }

    handleDragLeave = () => {
        this.dragging = false
    }

}

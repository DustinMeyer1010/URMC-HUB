export const allowedExtensions: string[] = [".txt", ".xlsx", ".csv"];

export class BulkLookUpStateClass {
    textValues: string = $state("");
    files: File[] = $state([]);
    loading: boolean = $state(false)

    updateFiles = (selectedFiles: File[]) => {
        this.files = selectedFiles
    }

    parseTextValueInput = () => {
        this.textValues = this.textValues.split("\n").map(line => line.trimStart().replaceAll("\t", " ")).join("\n")
    }

    lookupUsers = async () => {
        this.loading = true

        if (this.files.length == 0) {
            return
        }

        const formData = new FormData();
        formData.append("file", this.files[0])
        console.log(formData)

        const res = await fetch("http://localhost:8000/api/users/bulk/lookup", {
            method: "POST",
            body: formData
        });

        const blob = await res.blob();

        const filename =
            res.headers.get("Content-Disposition")?.split("filename=")[1]
            ?? "example.xlsx";

        const url = URL.createObjectURL(blob);

        const a = document.createElement("a");
        a.href = url;
        a.download = filename.replaceAll('"', '');
        a.style.display = "none";
        document.body.appendChild(a);
        a.click();            
        a.remove();
        URL.revokeObjectURL(url);

        this.loading = false
    }
    
}
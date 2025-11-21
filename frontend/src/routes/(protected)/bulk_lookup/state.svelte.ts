

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

        if (this.files.length != 0) {
            await this.fileLookup()
        } 
        else if (this.textValues != "") {
            await this.textLookup()
        } 

        this.loading = false
    }

    fileLookup = async () => {
        const formData = new FormData();
        formData.append("file", this.files[0])

        const res = await fetch("http://localhost:8000/api/users/bulk/lookup/file", {
            method: "POST",
            body: formData
        });

        await this.downloadFile(res)

    }

    textLookup = async () => {
        let values = this.textValues.split("\n")



        const res = await fetch("http://localhost:8000/api/users/bulk/lookup", {
            method: "POST",
            body: JSON.stringify(values)
        })

        await this.downloadFile(res)
    }

    async downloadFile(res: Response) {
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
    }
    
}
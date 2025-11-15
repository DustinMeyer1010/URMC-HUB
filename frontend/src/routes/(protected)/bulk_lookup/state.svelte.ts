export const allowedExtensions: string[] = [".txt", ".xlsx", ".csv"];

export class BulkLookUpStateClass {
    textValues: string = $state("");
    files: File[] = $state([]);

    updateFiles(selectedFiles: File[]) {
        this.files = selectedFiles
    }

    parseTextValueInput() {
        this.textValues = this.textValues.split("\n").map(line => line.trimStart().replaceAll("\t", " ")).join("\n")
    }
    
}
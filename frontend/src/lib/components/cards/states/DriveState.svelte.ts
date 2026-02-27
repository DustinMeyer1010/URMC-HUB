import { Drive } from "$lib/types/drive"
import { Copy } from "$lib/types/copy"

interface DriveState {
    groups: string[];
    drive: string;
    local_path: string;
    pageLink: string
    filteredGroups: string[];
    copyTemplate: string;
}

export class DriveStateClass implements DriveState {
    groups: string[] = []
    drive: string = ""
    local_path: string = ""
    searchValue: string = $state("")

    pageLink: string = $derived(`/drive?name=${encodeURIComponent(this.drive)}`)

    filteredGroups: string[] = $derived.by(() => {
        if (!this.searchValue) {
            return this.groups
        }


        return this.groups.filter((group) => {
            return group.toUpperCase().includes(this.searchValue.toUpperCase())
        })
    })

    copyTemplate: string = $derived.by(() => {
        const name: string = `Name = ${this.drive}`
        const localPath: string = `Local_Path = ${this.local_path}`
        const groups: string = `Groups: \n\n${this.filteredGroups.join("\n")}`
        return `${name}\n${localPath}\n\n${groups}`
    })

    constructor(drive: Drive.CardInfo) {
        Object.assign(this, drive)
    }

};
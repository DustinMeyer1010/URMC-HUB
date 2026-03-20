import { Drive } from "$lib/types/drive"

interface DriveState {
    groups: string[];
    drive: string;
    local_path: string;
    pageLink: string
    filteredGroups: string[];
    copyTemplate: string;
    groupCount: string;
}

export class DriveStateClass implements DriveState {
    groups: string[] = []
    drive: string = ""
    local_path: string = ""
    searchValue: string = $state("")

    pageLink: string = $derived(`/normal/drive?name=${encodeURIComponent(this.drive)}`)


    filteredGroups: string[] = $derived.by(() => {
        if (!this.searchValue) {
            return this.groups
        }


        return this.groups.filter((group) => {
            return group.toUpperCase().includes(this.searchValue.toUpperCase())
        })
    })

    groupCount: string = $derived(`Groups (${this.filteredGroups.length})`)

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
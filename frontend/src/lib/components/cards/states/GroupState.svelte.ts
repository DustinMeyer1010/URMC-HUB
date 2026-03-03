import { readableOU } from "$lib/parsers/ou";
import { Group } from "$lib/types/group";


interface GroupState {
    name: string
    information: string
    description: string
    ou: string
    readableOU: string
    pageLink: string
    copyTemplate: string
}


// * Lays out the entire group card and all the states that it will have
export class GroupStateClass implements GroupState {
    name: string = ""
    information: string = ""
    description: string = ""
    ou: string = ""

    readableOU: string = $derived(readableOU(this.ou))


    pageLink: string = $derived.by(() => {
        const name = encodeURIComponent(this.name)
        const ou = encodeURIComponent(this.ou)
        return `/group?name=${name}&ou=${ou}`
    })

    copyTemplate: string = $derived.by(() => {
        const info = this.information || "NA"
        const desc = this.description || "NA"
        return `Name = ${this.name}\nInformation = ${info}\nDescription = ${desc}\nOU = ${this.readableOU}`
    })

    constructor(group: Group.CardInfo) {
        Object.assign(this, group)
    }


}


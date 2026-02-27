import type { User } from "$lib/types/user";
import { readableOU } from "$lib/utils/stringEditor";

interface UserState {
    name: string
    username: string
    net_id: string
    urid: string
    email: string
    ou: string
    pageLink: string
    readableOU: string
    disabled: boolean
    copyTemplate: string
}

export class UserStateClass implements UserState {
    name = "";
    username = "";
    net_id = "";
    urid = "";
    email = "";
    ou = "";

    readableOU: string = $derived(readableOU(this.ou))
    
    disabled: boolean = $derived.by(() => {
        const ou = this.ou.toLowerCase()
        return ou.includes("disabled") || ou.includes("offboarded")
    })

    // TODO: Turn in query for the user and OU rather than just searching by person username
    pageLink: string = $derived.by(() => {
      return `/user/${this.username}`  
    })

    copyTemplate: string = $derived.by(() => {
        const name: string = `Name = ${this.name || "NA"}`
        const username: string = `Username = ${this.username || "NA"}`
        const netID: string = `NetID = ${this.net_id || "NA"}`
        const urid: string = `URID = ${this.urid || "NA"}`
        const email: string = `Email = ${this.email || "NA"}`

        return `${name}\n${username}\n${netID}\n${urid}\n${email}\n${this.readableOU}\n`
    })


    constructor(user: User.CardInfo) {
        Object.assign(this, user)
    }
}
import { readableDN } from "$lib/parsers/ou";
import type { User } from "$lib/types/user";

interface UserState {
    cn: string
    username: string
    netid: string
    urid: string
    email: string
    dn: string
    pageLink: string
    readableDN: string
    disabled: boolean
    copyTemplate: string
}

export class UserStateClass implements UserState {
    cn = "";
    username = "";
    netid = "";
    urid = "";
    email = "";
    dn = "";

    readableDN: string = $derived(readableDN(this.dn))
    
    disabled: boolean = $derived.by(() => {
        const dn = this.dn.toLowerCase()
        return dn.includes("disabled") || dn.includes("offboarded")
    })

    // TODO: Turn in query for the user and OU rather than just searching by person username
    pageLink: string = $derived.by(() => {
      return `/user?dn=${encodeURIComponent(this.dn)}&section=PROFILE`  
    })

    copyTemplate: string = $derived.by(() => {
        const name: string = `Name = ${this.cn || "NA"}`
        const username: string = `Username = ${this.username || "NA"}`
        const netID: string = `NetID = ${this.netid || "NA"}`
        const urid: string = `URID = ${this.urid || "NA"}`
        const email: string = `Email = ${this.email || "NA"}`

        return `${name}\n${username}\n${netID}\n${urid}\n${email}\n${this.readableDN}\n`
    })

    simpleCopyTemplate: string = $derived.by(() => {

        return `${this.cn} (${this.username})`
    })


    constructor(user: User.CardInfo) {

        console.log(user)
        this.cn = user.cn.join()
        this.username = user.username.join()
        this.netid = user.netid.join()
        this.urid = user.urid.join()
        this.email = user.email.join()
        this.dn = user.dn.join()
    }
}

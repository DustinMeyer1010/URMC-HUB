
import { readableDN } from "$lib/parsers/ou";
import {Computer} from "$lib/types/computer";

interface ComputerState {
    name: string
    ou: string
    operating_system: string

    disabled: boolean
    readableOU: string
    copyTemplate: string
    pageLink: string
}

export class ComputerStateClass implements ComputerState {
    name: string = ""
    ou: string = ""
    operating_system: string = ""

    disabled: boolean = $derived(this.ou.toLowerCase().includes("disabled"))

    readableOU: string = $derived(readableDN(this.ou))

    pageLink: string = $derived.by(() => {
        return `/computer/${this.name}`
    })


    copyTemplate: string = $derived.by(() => {
        const osSuffix = this.operating_system ? ` (${this.operating_system})` : ""
        const status = this.disabled ? "Disabled " : ""
        return `Hostname = ${this.name}${osSuffix}\n${status}OU = ${this.readableOU}`
    })


    constructor(computer: Computer.CardInfo) {
        Object.assign(this, computer)
    }
}

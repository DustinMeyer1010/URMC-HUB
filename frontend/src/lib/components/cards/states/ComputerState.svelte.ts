
import { readableDN } from "$lib/parsers/ou";
import {Computer} from "$lib/types/computer";


export class ComputerStateClass  {
    cn: string = ""
    dn: string = ""
    operatingSystem: string = ""
    information: string = ""
    description: string = ""

    disabled: boolean = $derived(this.dn.toLowerCase().includes("disabled"))

    readableOU: string = $derived(readableDN(this.dn))

    pageLink: string = $derived.by(() => {
        return `/normal/computer?dn=${this.dn}`
    })


    copyTemplate: string = $derived.by(() => {
        const osSuffix = this.operatingSystem ? ` (${this.operatingSystem})` : ""
        const description = this.description ? `Description = ${this.description}` : ""
        const status = this.disabled ? "Disabled " : ""
        return `Hostname = ${this.cn}${osSuffix}\n${status}${description}\nOU = ${this.readableOU}`
    })


    constructor(computer: Computer.CardInfo) {
        this.cn = computer.cn.join()
        this.dn = computer.dn.join()
        this.operatingSystem = computer.operatingsystem.join()
        this.information = computer.information.join()
        this.description = computer.description.join()
    }
}

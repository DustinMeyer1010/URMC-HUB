import { readableDN } from "$lib/parsers/ou";
import { Group } from "$lib/types/group";



// * Lays out the entire group card and all the states that it will have
export class GroupStateClass  {
    cn: string = ""
    samaccountname: string = ""
    information: string = ""
    dn: string = ""
    description: string = ""

    readableOU: string = $derived(readableDN(this.dn))

    sigGroup: boolean = $derived(this.dn.toLocaleLowerCase().includes("sig"))


    pageLink: string = $derived.by(() => {
        const name = encodeURIComponent(this.cn)
        const dn = encodeURIComponent(this.dn)
        return `/normal/group?name=${name}&ou=${dn}`
    })

    copyTemplate: string = $derived.by(() => {
        const info = this.information || "NA"
        const desc = this.description || "NA"
        return `Name = ${this.cn}\nInformation = ${info}\nDescription = ${desc}\nOU = ${this.readableOU}`
    })

    constructor(group: Group.CardInfo) {
        console.log(group)
        this.cn = group.cn.join()
        this.samaccountname = group.samaccountname.join()
        this.information = group.information.join()
        this.description = group.description.join()
        this.dn = group.dn.join()
    }


}


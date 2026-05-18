import { readableDN } from "$lib/parsers/ou";
import type { User } from "$lib/types/user";


export class UserStateClass {
    cn: string = "";
    username: string = "";
    netid: string = "";
    urid: string = "";
    email: string = "";
    dn: string = "";
    memberof: string[] = [];
    passwordLastSet = "0";

    readableDN: string = $derived(readableDN(this.dn))
    
    disabled: boolean = $derived.by(() => {
        const dn = this.dn.toLowerCase()
        return dn.includes("disabled") || dn.includes("offboarded")
    })

    idleAccount: boolean = $derived.by(() => {
        for (let i=0; i < this.memberof.length; i++){
            if (this.memberof[i].toLowerCase().includes("idm_idleaccounts_urmc")){
                return true
            }
        }
        return false
    })

    // Refactor: This will need to be seperated into a function that does the conversions
    expiredPassword: boolean = $derived.by(() => {
        const val = BigInt(this.passwordLastSet);
        if (val === 0n) return true;
        const windowsEpochOffset = 116444736000000000n;
        const millisecondsSinceEpoch = (val - windowsEpochOffset) / 10000n;

        const passwordDate = new Date(Number(millisecondsSinceEpoch));
        const oneYearAgo = new Date();
        oneYearAgo.setFullYear(oneYearAgo.getFullYear() - 1);

        return passwordDate < oneYearAgo;
    })

    pageLink: string = $derived.by(() => {
      return `/normal/user?dn=${encodeURIComponent(this.dn)}&section=GROUPS`  
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
        this.cn = user.cn.join()
        this.username = user.username.join()
        this.netid = user.netid.join()
        this.urid = user.urid.join()
        this.email = user.email.join()
        this.dn = user.dn.join()
        this.memberof = user.memberof
        this.passwordLastSet = user.pwdlastset.join()
    }
}

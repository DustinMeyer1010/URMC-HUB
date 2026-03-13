import { type Drive } from "./drive"


export namespace User {

    export type Section = "PROFILE" | "LOCKOUT" | "DRIVES" | "GROUPS" | "ADD"
    export const Sections: string[] = ["PROFILE", "LOCKOUT", "DRIVES", "GROUPS", "ADD"]

    export const isValidSection = (value: string): boolean => {
        return Sections.includes(value.toUpperCase())
    }

    export type CardInfo = {
        cn: string[],
        username: string[], 
        netid: string[],
        urid: string[],
        email: string[],
        dn: string[],
    }

    
    export type PageInfo = {
        username: string,
        name: string, 
        email: string, 
        relationship_status: string,
        department: string,
        title: string,
        phone: string,
        location: string,
        last_password_set: string,
        urid: string,
        ou: string,
        net_id: string,
        description: string,
        member_of: Drive.CardInfo[],
    }

    export type LockoutInfo = {
        name: string,
        count: number,
        time: string,   
    }




}





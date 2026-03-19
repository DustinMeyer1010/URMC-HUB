export namespace Group {
    export const EMPTY_GROUP: CardInfo = {
        cn: [],
        information: [],
        description: [],
        dn: [],
        samaccountname: []
    } 
    export type Section = "PROFILE" | "LOCKOUT" | "DRIVES" | "GROUPS" | "ADD" | "REMOVE"
    export const Sections: string[] = ["PROFILE", "LOCKOUT", "DRIVES", "GROUPS", "ADD", "REMOVE"]

    export const isValidSection = (value: string): boolean => {
        return Sections.includes(value.toUpperCase())
    }

    export type CardInfo = {
        cn: string[],
        samaccountname: string[],
        information: string[],
        description: string[],
        dn: string[],
    }

    export type PageInfo = {
        name: string,
        ou: string,
        information: string,
        description: string,
        memebers: string
    }
}
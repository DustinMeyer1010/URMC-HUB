export namespace Group {
    export const EMPTY_GROUP: CardInfo = {
        name: "",
        information: "",
        description: "",
        ou: ""
    } 
    export type Section = "PROFILE" | "LOCKOUT" | "DRIVES" | "GROUPS" | "ADD" | "REMOVE"
    export const Sections: string[] = ["PROFILE", "LOCKOUT", "DRIVES", "GROUPS", "ADD", "REMOVE"]

    export const isValidSection = (value: string): boolean => {
        return Sections.includes(value.toUpperCase())
    }

    export type CardInfo = {
        name: string,
        information: string,
        description: string,
        ou: string,
    }

    export type PageInfo = {
        name: string,
        ou: string,
        information: string,
        description: string,
        memebers: string
    }
}
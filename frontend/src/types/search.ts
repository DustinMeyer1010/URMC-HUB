import type { Computer } from "./computer"
import type { Printer } from "./printer"
import type { User } from "./user"
import type { Group } from "./group"
import type { Drive } from "./drive"

export namespace Search {
    export type Filters = 'COMPUTERS' | 'USERS' | "GROUPS" | 'PRINTERS' | 'DRIVES'
    export const ValidFilters: string[] = ['COMPUTERS' , 'USERS' , "GROUPS" , 'PRINTERS' , 'DRIVES']  

    export type Results = {
        users: User.CardInfo[],
        computers: Computer.CardInfo[],
        groups: Group.CardInfo[],
        printers: Printer.CardInfo[],
        drives: Drive.CardInfo[],
    }

    export const isValidFilter = (value: string): boolean => {
        return ValidFilters.includes(value.toUpperCase())
    }
}
import type { ComputerSimpleInfo } from "./computer"
import type { PrinterSimpleInfo } from "./printer"
import type { UserSimpleInfo } from "./user"
import type { GroupSimpleInfo } from "./group"
import type { DriveSimpleInfo } from "./drive"


export type Filters = 'COMPUTERS' | 'USERS' | "GROUPS" | 'PRINTERS' | 'DRIVES'
export const AllFilters: string[] = ['COMPUTERS' , 'USERS' , "GROUPS" , 'PRINTERS' , 'DRIVES']
export type Results =  ComputerSimpleInfo[] | PrinterSimpleInfo[] | UserSimpleInfo[] | GroupSimpleInfo[] | DriveSimpleInfo[]



export namespace Filter {
    export const isValid = (value: string): boolean => {
        return AllFilters.includes(value);
    }
}
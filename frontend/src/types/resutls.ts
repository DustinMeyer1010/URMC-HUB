import type { ComputerSimpleInfo } from "./computer"
import type { GroupSimpleInfo } from "./group"
import type { PrinterSimpleInfo } from "./printer"
import type { UserSimpleInfo } from "./user"
import type { DriveSimpleInfo } from "./drive"



export type AllResults = {
    users: UserSimpleInfo[],
    computers: ComputerSimpleInfo[],
    groups: GroupSimpleInfo[],
    printers: PrinterSimpleInfo[],
    drives: DriveSimpleInfo[],
}



export type ModifyResults = {
    group: string,
    message: string,
    successful: boolean,
}

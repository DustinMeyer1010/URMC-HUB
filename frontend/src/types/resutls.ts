import type { Computer } from "./computer"
import type { GroupSimpleInfo } from "./group"
import type { PrinterSimpleInfo } from "./printer"
import type { User } from "./user"
import type { Drive } from "./drive"




export namespace Modification {
    export type Results = {
        group: string,
        message: string,
        successful: boolean,
    }
}

import type { Computer } from "./computer"
import type { Printer } from "./printer"
import type { User } from "./user"
import type { Group } from "./group"
import type { Drive } from "./drive"




export type Results = Computer.CardInfo[] | Printer.CardInfo[] | User.CardInfo[] | Group.CardInfo[] | Drive.CardInfo[]

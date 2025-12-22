import { Drive } from "@t/drive"
import { Copy } from "@t/copy"

export class DriveStateClass {
    SearchValue: string = $state("")
    Idx: number = $state(0)
    CurrentDrive: Drive.CardInfo = $state(Drive.EMPTY_DRIVE_CARD)
    CopyState: Copy.State = $state(Copy.EMPTY_COPY_STATE)

    Groups: string[] = $derived.by(() => {
        if (!this.SearchValue) {
            return this.CurrentDrive.groups
        }
        return this.CurrentDrive.groups.filter((group) => group.toUpperCase().includes(this.SearchValue.toUpperCase()))
    })

    SerializedSearch: string = $derived(encodeURIComponent(this.CurrentDrive.drive))

    CopyText: string = $derived(`Name: ${this.CurrentDrive.drive}\nLocal_Path: ${this.CurrentDrive.local_path}\nGroups:\n${this.CurrentDrive.groups.join("\n")}`);
    
    constructor(idx: number, currentDrive: Drive.CardInfo) {
        this.Idx = idx
        this.CurrentDrive = currentDrive
    }

}
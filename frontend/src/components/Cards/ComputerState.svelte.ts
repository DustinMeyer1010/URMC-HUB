import { Copy } from "@t/copy"

interface ComputerState {
    CopyState: Copy.State
    Disabled: boolean
    AllText: string
}

export class ComputerStateClass implements ComputerState {
    CopyState: Copy.State = $state(
        {
            copied: "",
            timeout: null
        }
    )

    constructor(disabled: boolean, allText: string) {
        this.Disabled = disabled
        this.AllText = allText
    }


    Disabled: boolean = $state(false)
    AllText: string = $state("")
}
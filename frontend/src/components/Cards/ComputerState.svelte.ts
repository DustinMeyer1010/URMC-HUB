import type { CopyState } from "$lib/helper/copy.svelte";

interface ComputerState {
    CopyState: CopyState
    Disabled: boolean
    AllText: string
}

export class ComputerStateClass implements ComputerState {
    CopyState: CopyState = $state(
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
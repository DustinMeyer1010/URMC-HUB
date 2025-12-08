import type { ComputerPageInfo } from "@t/computer"

export const EMPTY_COMPUTER_INFO: ComputerPageInfo = {
    computer_info: {
        name: "",
        ou: "",
        operating_system: ""
    },
    is_online: false,
    ping_results: ""
}

interface PageState {
    ComputerName: string
    ComputerInformation: ComputerPageInfo
    Disabled: boolean
    Loading: boolean
    GetComputerInfo: () => void
}


export class PageStateClass implements PageState {
    Loading: boolean = $state(true)
    ComputerInformation: ComputerPageInfo = $state(EMPTY_COMPUTER_INFO)
    Disabled: boolean = $derived.by(() => this.ComputerInformation.computer_info.ou.toUpperCase().includes("DISABLED"))
    ComputerName: string;

    constructor(computer: string) {
        this.ComputerName = computer
    }

    GetComputerInfo = async () => {
        this.Loading = true
        await fetch(`http://localhost:8000/api/computer/${this.ComputerName}/info`)
        .then(async (res) => {
            this.ComputerInformation = await res.json()
        })
        this.Loading = false
    }
}
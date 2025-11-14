

export type ComputerSimpleInfo = {
    name: string,
    ou: string,
    operating_system: string
}   




export type ComputerPageInfo = {
    computer_info: ComputerSimpleInfo,
    is_online: boolean
    ping_results: string
}
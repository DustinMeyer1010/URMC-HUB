export namespace Computer {
    export type CardInfo = {
        name: string,
        ou: string,
        operating_system: string
    }

    export type PageInfo = {
        computer_info: CardInfo,
        is_online: boolean
        ping_results: string
    }
}

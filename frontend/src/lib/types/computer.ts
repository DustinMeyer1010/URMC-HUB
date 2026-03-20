export namespace Computer {

    export const EMPTY_CARDINFO: CardInfo = {
        cn: [],
        dn: [],
        operatingsystem: [],
        information: [],
        description: [],
    }

    export type CardInfo = {
        cn: string[],
        dn: string[],
        operatingsystem: string[],
        information: string[],
        description: string[],
    }

    export type PageInfo = {
        computer_info: CardInfo,
        is_online: boolean
        ping_results: string
    }
}

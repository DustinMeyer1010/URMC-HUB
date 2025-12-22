export namespace Drive {
    export const EMPTY_DRIVE_CARD: CardInfo = {
        groups: [],
        drive: "",
        local_path: ""
    }

    export type CardInfo = {
        groups: string[],
        drive: string,
        local_path: string,
    }
}
export namespace Printer {
    export const EMPTY_PRINTER: CardInfo = {
        server: "",
        queue: "",
        model: "",
        ip: "",
        print_processor: "",
        location: "",
        notes: ""
    }

    export type CardInfo = {
        server: string,
        queue: string, 
        model: string,
        ip: string,
        print_processor: string,
        location: string, 
        notes: string
    }
}


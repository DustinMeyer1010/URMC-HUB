/*
type PrinterSimpleInfo struct {
	Server         string `json:"server"`
	Queue          string `json:"queue"`
	Model          string `json:"model"`
	IP             string `json:"ip"`
	PrintProcessor string `json:"print_processor"`
	Location       string `json:"location"`
	Notes          string `json:"notes"`
}
 */


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


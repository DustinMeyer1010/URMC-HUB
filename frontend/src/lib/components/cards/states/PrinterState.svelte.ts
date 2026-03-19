import { Printer } from "$lib/types/printer";



interface PrinterState {
    server: string;
    queue: string;
    model: string;
    ip: string;
    print_processor: string;
    location: string;
    notes: string;
    fullName: string;

    pageLink: string;
    copyTemplate: string
}

export class PrinterStateClass implements PrinterState {
    server: string = "";
    queue: string  = "";
    model: string = "";
    ip: string = "";
    print_processor: string = "";
    location: string = "";
    notes: string = "";

    readableNotes: string = $derived(this.notes.replace(/[,\s]+/g, ",").replace(/,$/, "").replaceAll(",", " - "))
    fullName: string = $derived(`\\\\${this.server}\\${this.queue}`)

    pageLink: string = $derived.by(() => {
        return `/printer?server=${this.server}&queue=${this.queue}`
    })

    copyTemplate: string = $derived.by(() => {
        const name = `Name = ${this.fullName}`
        const model = `Model = ${this.model || "NA"}`
        const ip = `Model = ${this.ip || "NA"}`
        const print_processor = `Processor = ${this.print_processor || "NA"}`
        const location = `Location = ${this.location || "NA"}`
        return `${name}\n${model}\n${ip}\n${print_processor}\n${location}\n${this.readableNotes}`
    })


    constructor(printer: Printer.CardInfo) {
        Object.assign(this, printer)
    }
}   




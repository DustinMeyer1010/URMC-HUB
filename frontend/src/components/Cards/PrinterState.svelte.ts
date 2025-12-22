import { Printer } from "@t/printer";

export class PrinterStateClass {
    Printer: Printer.CardInfo = Printer.EMPTY_PRINTER
    URLQuery: string = $derived(`${this.Printer.server}?queue=${this.Printer.queue}`)
    CopyText: string =  $derived(`Name: \\\\${this.Printer.server}\\${this.Printer.queue}\nModel: ${this.Printer.model}\nIP: ${this.Printer.ip}\nPrint_Processor: ${this.Printer.print_processor}\nLocation: ${this.Printer.location}\nNotes: ${this.Printer.notes}`)
    PrinterName: string = $derived(`\\\\${this.Printer.server}\\${this.Printer.queue}`)
    Idx: number = $state(0)

    constructor(printer: Printer.CardInfo, idx: number) {
        this.Printer = printer
        this.Idx = idx
    }
}   
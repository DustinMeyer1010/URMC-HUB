export type PrinterCardComponent = {
    Name?: string,
    Model?: string
    IP?: string,
    Processor?: string,
    Location?: string,
    Notes?: string
    Select: (Type: string, Name: string) => void,
    ItemsSelected: string[],

}
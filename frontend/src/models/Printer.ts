export type PrinterCardComponent = {
    name?: string,
    model?: string
    ip?: string,
    processor?: string,
    location?: string,
    notes?: string
    Select: (Type: string, Name: string) => void,
    ItemsSelected: string[],

}
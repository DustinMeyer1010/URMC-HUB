export type ComputerCardInfo = {
    name?: string;
    ou?: string;
    operating_system?: string;
}

export type ComputerCardComponent = {
    name?: string;
    ou?: string;
    operating_system?: string;
    Select: (Type: string, Name: string) => void,
    ItemsSelected: string[],
}


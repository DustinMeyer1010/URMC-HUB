export type ComputerCardInfo = {
    Name?: string;
    OU?: string;
    OperatingSystem?: string;
}

export type ComputerCardComponent = {
    Name?: string;
    OU?: string;
    OperatingSystem?: string;
    Select: (Type: string, Name: string) => void
}
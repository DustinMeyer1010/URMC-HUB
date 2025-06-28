export type GroupCardInfo = {
    Type?: string,
    OU?: string,
    Name?: string,
    Description?: string,
    Information?: string
}

export type GroupCardComponent = {
    Name?: string;
    OU?: string;
    Description?: string,
    Information?: string, 
    Select: (Type: string, Name: string) => void,
    ItemsSelected: string[],
}
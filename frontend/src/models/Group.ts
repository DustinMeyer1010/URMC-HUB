export type GroupCardInfo = {
    name?: string,
    ou?: string,
    description?: string,
    information?: string
}

export type GroupCardComponent = {
    name?: string;
    ou?: string;
    description?: string,
    information?: string, 
    Select: (Type: string, Name: string) => void,
    ItemsSelected: string[],
}
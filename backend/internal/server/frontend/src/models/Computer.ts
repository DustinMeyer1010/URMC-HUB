export type ComputerCardInfo = {
    Name?: string;
    OU?: string;
    OperatingSystem?: string;
}

export type ComputerCardComponent = {
    Name?: string;
    OU?: string;
    OperatingSystem?: string;
    Select: (Type: string, Name: string) => void,
    ItemsSelected: string[],
}

export const CopyComputerAtrribute = (text: string) => {
    navigator.clipboard.writeText(text)
    .then(() => {

    })
    .catch((err) => {
        alert(err)
    })
}   
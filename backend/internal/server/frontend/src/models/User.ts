export type UserCardInfo = {
    Type?: string;
    Name?: string,
    Username?: string,
    NetID?: string,
    Email?: string,
    OU?: string,
}


export type UserCardComponent = {
    Name?: string,
    Username?: string,
    NetID?: string,
    Email?: string,
    OU?: string,
    Select: (Type: string, Name: string) => void,
    ItemsSelected: string[],
}

export const CopyUserAtrribute = (text: string) => {
    navigator.clipboard.writeText(text)
    .then(() => {

    })
    .catch((err) => {
        alert(err)
    })
}   
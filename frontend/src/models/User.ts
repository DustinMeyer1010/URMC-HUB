export type UserCardInfo = {
    name?: string,
    username?: string,
    net_id?: string,
    email?: string,
    ou?: string,
}


export type UserCardComponent = {
    name?: string,
    username?: string,
    net_id?: string,
    email?: string,
    ou?: string,
    Select: (Type: string, Name: string) => void,
    ItemsSelected: string[],
}



export type UserInformation = {
    name?: string,
    username?: string, 
    net_id?: string, 
    email?: string,
    ou?: string,
    relationship_status: string,
    department?: string, 
    title?: string,
    phone?: string,
    location?: string,
    last_password_set?: string,
    urid?: string, 
    descirption?: string,
}
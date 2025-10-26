
export type UserSimpleInfo = {
    name: string,
    username: string, 
    net_id: string,
    urid: string,
    email: string,
    ou: string,
}


export type UserFullInfo = {
    username: string,
    name: string, 
    email: string, 
    relationship_status: string,
    department: string,
    title: string,
    phone: string,
    location: string,
    last_password_set: string,
    urid: string,
    ou: string,
    net_id: string,
    description: string,
    member_of: string[],
}


export type LockoutInfo = {
    name: string,
    count: number,
    time: string,
}
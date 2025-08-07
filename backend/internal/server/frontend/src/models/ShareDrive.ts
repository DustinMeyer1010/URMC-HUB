export type ShareDriveComponent = {
    Server: string,
    Name: string,
    Groups: string[],
    Select: (Type: string, Name: string) => void,
    ItemsSelected: string[],
}
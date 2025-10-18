export type ShareDriveComponent = {
    server: string,
    name: string,
    groups: string[],
    Select: (Type: string, Name: string) => void,
    ItemsSelected: string[],
}
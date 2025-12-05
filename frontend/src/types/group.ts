/**
 * For Cards to display infromation about a group
*/ 
export type GroupSimpleInfo = {
    name: string,
    ou: string,
    information: string,
    description: string,
}
/**
 * Informatino for displaying on group page
 */
export type GroupPageInfo = {
    name: string,
    ou: string,
    information: string,
    description: string,
    memebers: string
}
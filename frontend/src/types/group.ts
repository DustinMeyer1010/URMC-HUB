/**
 * For Cards to display infromation about a group
*/ 
export type GroupSimpleInfo = {
    name: string,
    information: string,
    description: string,
    ou: string,
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
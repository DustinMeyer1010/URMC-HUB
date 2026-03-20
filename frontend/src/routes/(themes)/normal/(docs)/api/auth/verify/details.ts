export const endpoint = `GET /api/user/verify HTTP/1.1`

export const response = `HTTP/1.1 200 OK\nHTTP/1.1 401 Unauthorized`


export const information = `
The Verify endpoint will verify that the user is still logged in. 
Checks to see if the credentials still work against the ldap server.
`


export const responses: Responses[] = [
    {
        status: "OK",
        code: 200,
        description: "Still authenticated"
    },
    {
        status: "Unauthorized",
        code: 401,
        description: "No longer authenticated"
    },
]
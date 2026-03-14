export const request = JSON.stringify(
    {
		username: "username",
		password: "user_password",
	}, 
    null, 4);

export const endpoint = `POST /api/user/login HTTP/1.1`

export const response = `HTTP/1.1 200 OK\nHTTP/1.1 400 Bad Request\nHTTP/1.1 401 Unauthorized`


export const information = `The login endpoint will login agent into application.
                Once the agent is logged in they will be able to access the protected routes.
                This will check the agents credentials against the ldap server. 
                Once there is a successful bind the agent is logged in.`


export const objects: Attributes[]  = [
    {
        name: "username", 
        type: "string" , 
        description: "Agents username"
    },
    {
        name: "password", 
        type: "string", 
        description: "Agents Password"
    },
]

export const responses: Responses[] = [
    {
        status: "OK",
        code: 200,
        description: "Login Successful"
    },
    {
        status: "Bad Request",
        code: 400,
        description: "Invalid body"
    },
    {
        status: "Unauthorized",
        code: 401,
        description: "Invalid username or password"
    },
]
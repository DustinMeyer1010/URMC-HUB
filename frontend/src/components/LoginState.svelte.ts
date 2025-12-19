type LoginForm = {
    username: string,
    password: string
}

let EMPTY_LOGIN_FORM: LoginForm = {
    username: "",
    password: ""
}

export class LoginStateClass {
    
    Error: boolean = $state(false)
    Form: LoginForm = $state(EMPTY_LOGIN_FORM)
    Login: () => void;

    constructor(login: () => void) {
        this.Login = login
    }

    onsubmit = async (e: SubmitEvent) => {
        e.preventDefault()
        await fetch("http://localhost:8000/api/user/login", {
            method: "POST",
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(this.Form)
        })
        .then(async (res) => {
            if (res.ok) {
                window.location.reload()
                localStorage.setItem("agent", this.Form.username)
                this.Login()
                return
            }
        })

        this.Error = true
    }

}
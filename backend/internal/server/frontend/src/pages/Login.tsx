import { useAuth } from "../components/Authentication";
import input from "../styles/Input.module.css"
import Button from "../styles/Button.module.css"
import Form from "../styles/Form.module.css"
import LoginStyles from "./Login.module.css"
import { useNavigate } from "react-router-dom";

const Login = () => {

    const { login } = useAuth()
    const navigate = useNavigate()

    const handleLogin = (e: React.FormEvent) => {
        e.preventDefault()
    
        login()
        navigate("/search")
    }

    return (
        <>
        <form onSubmit={handleLogin} className={Form.default} action="">
            <h1>Login</h1>
            <div className={LoginStyles.input_container}>
                <input placeholder=" " type="text" className={`${input.default} ${LoginStyles.input}`}/>
                <label className={LoginStyles.label}>Username</label>
            </div>
            <div className={LoginStyles.input_container}>
                <input placeholder=" " type="password" className={`${input.default} ${LoginStyles.input}`}/>
                <label className={LoginStyles.label}>Password</label>
            </div>
                <button type="submit" className={`${Button.default} ${LoginStyles.button_login}`}>Login</button>
        </form>
        </>

    )
}

export default Login;
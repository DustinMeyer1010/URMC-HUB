import { useAuth } from "../components/Authentication";

const Login = () => {

    const { login } = useAuth()
    

    const handleLogin = () => {
        login()
    }

    return (
        <>
        <button onClick={handleLogin}>Login</button>
        </>

    )
}

export default Login;
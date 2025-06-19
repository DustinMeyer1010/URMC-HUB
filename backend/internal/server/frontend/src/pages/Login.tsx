import { useNavigate } from "react-router-dom";
import { useState } from "react";
import axios from "axios";

import { useAuth } from "../components/Authentication";

import input from "../styles/global/Input.module.css"
import Button from "../styles/global/Button.module.css"
import Form from "../styles/global/Form.module.css"
import LoginStyles from "../styles/pages/Login.module.css"


const Login = () => {
    type LoginForm = {
        username: string;
        password: string;
    }

    const [formData, setFormData] = useState<LoginForm>({
        username: '', 
        password: '',
    })

    const { login } = useAuth()
    const navigate = useNavigate()

    const handleLogin = async (e: React.FormEvent) => {
        e.preventDefault()

        try {
            const response = await axios.post("http://localhost:8080/user/login", formData)

            if (response.status == 200){
                login()
                navigate("/search")
            }
        } catch (error) {

            console.log(error)
        }
    }

    const handleValueChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const {name, value} = e.target;

        setFormData((prev) => ({
            ...prev,
            [name]: value,
        }))

    }

    return (
        <form onSubmit={handleLogin} className={Form.default} action="">
            <h1>Login</h1>
            <div className={LoginStyles.input_container}>
                <input 
                    placeholder=" "
                    name="username" 
                    onChange={handleValueChange}
                    type="text" 
                    value={formData.username} 
                    className={`${input.default} ${LoginStyles.input}`}
                />
                <label className={LoginStyles.label}>Username</label>
            </div>
            <div className={LoginStyles.input_container}>
                <input 
                    placeholder=" "
                    name="password" 
                    onChange={handleValueChange} 
                    type="password" 
                    value={formData.password}
                    className={`${input.default} ${LoginStyles.input}`}
                />
                <label className={LoginStyles.label}>Password</label>
            </div>
                <button type="submit" className={`${Button.default} ${LoginStyles.button_login}`}>Login</button>
        </form>

    )
}

export default Login;
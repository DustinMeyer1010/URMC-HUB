import { useParams } from "react-router-dom"

const User = () => {
    const { Username } = useParams<{ Username: string }>();

    return (
        <div>
            <h1>Hello { Username } Page</h1>
        </div>
    )
}


export default User
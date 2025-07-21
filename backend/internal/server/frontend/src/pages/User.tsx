import { useParams } from "react-router-dom"
import Styles from "../styles/pages/User.module.css"
import React from "react";

const User = () => {
    const { Username } = useParams<{ Username: string }>();

    React.useEffect(() => {
        document.body.classList.toggle("no-scroll")

        return () => {
            document.body.classList.remove("no-scroll")
        }
    }) 

    const handleContentChange = () => {
        document.getElementById("info")?.classList.toggle(Styles.move)
         document.getElementById("info2")?.classList.toggle(Styles.move1)

    }

    return (
        <div className={Styles.layout}>
            <h1>User Page</h1>
            <div className={Styles.userInfo}>
                <button onClick={handleContentChange}>Test</button>
                <ul id="info" className={Styles.content}>
                    <li>Username: {Username}</li>
                    <li>NedID</li>
                    <li>Email</li>
                    <li>Relationship Status</li>
                    <li>Department</li>
                    <li>Title</li>
                    <li>Phone</li>
                    <li>Location</li>
                    <li>Password Set</li>
                </ul>
                <ul id="info1" className={`${Styles.content} ${Styles.move1}`}>
                    <li>Username: {Username}</li>
                    <li>NedID</li>
                    <li>Email</li>
                    <li>Relationship Status</li>
                    <li>Department</li>
                    <li>Title</li>
                    <li>Phone</li>
                    <li>Location</li>
                    <li>Password Set</li>
                </ul>
                <table id="info2" className={`${Styles.content} ${Styles.move1}`}>
                    <thead>
                        <th>Name</th>
                        <th>Count</th>
                        <th>Time</th>
                    </thead>
                </table>
            </div>
            
        </div>
    )
}


export default User
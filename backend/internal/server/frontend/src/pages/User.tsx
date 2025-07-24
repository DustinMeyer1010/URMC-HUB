import { useParams } from "react-router-dom"
import Styles from "../styles/pages/User.module.css"
import React, { useState } from "react";

const User = () => {
    const { Username } = useParams<{ Username: string }>();
    const [currentContent, setCurrentContent] = useState({ id: "info", move: Styles.moveInfo})

    React.useEffect(() => {
        document.body.classList.toggle("no-scroll")

        return () => {
            document.body.classList.remove("no-scroll")
        }
    }) 

    const handleContentChange = (id: string, move: string) => {
        if (currentContent.id == id){
            return
        }

        document.getElementById(currentContent.id)?.classList.toggle(currentContent.move)

        document.getElementById(id)?.classList.toggle(move)

        setCurrentContent({id: id, move: move})
    }

    return (
        <div className={Styles.layout}>
            <h1>Dustin Meyer</h1>
            <div className={Styles.userInfo}>
                <div className={Styles.navigation}>
                    <button onClick={() => handleContentChange("info", `${Styles.moveInfo}`)}>About</button>
                    <button onClick={() => handleContentChange("groups", `${Styles.moveGroups}`)}>Groups</button>
                    <button onClick={() => handleContentChange("sharedrive", `${Styles.moveShareDrive}`)}>Share Drive</button>
                    <button onClick={() => handleContentChange("lockout", `${Styles.moveLockout}`)}>Lockout</button>
                </div>

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
                <ul id="groups" className={`${Styles.content} ${Styles.moveGroups}`}>
                    Groups content
                </ul>
                <ul id="sharedrive" className={`${Styles.content} ${Styles.moveShareDrive}`}>
                    Sharedrive content
                </ul>
                <table id="lockout" className={`${Styles.content} ${Styles.table} ${Styles.moveLockout}`}>
                    <colgroup>
                        <col style={{width: "33%"}}/>
                        <col style={{width: "33%"}}/>
                        <col style={{width: "33%"}}/>
                    </colgroup>
                    <thead>
                        <tr>
                            <th>Name</th>
                            <th>Count</th>
                            <th>Time</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td>AD22PDC01</td>
                            <td>0</td>
                            <td>07/11/2025 16:59:55</td>
                        </tr>
                                                <tr>
                            <td>AD22PDC01</td>
                            <td>0</td>
                            <td>07/11/2025 16:59:55</td>
                        </tr>
                                                <tr>
                            <td>AD22PDC01</td>
                            <td>0</td>
                            <td>07/11/2025 16:59:55</td>
                        </tr>
                        <tr>
                            <td>AD22PDC01</td>
                            <td>0</td>
                            <td>07/11/2025 16:59:55</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    )
}


export default User
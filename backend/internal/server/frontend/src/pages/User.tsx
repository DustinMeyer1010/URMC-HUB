import { useParams } from "react-router-dom"
import Styles from "../styles/pages/User.module.css"
import React, { useState } from "react";
import axios from "axios";
import type { UserInformation }from "../models/User"

const User = () => {
    const { Username } = useParams<{ Username: string }>();
    const [currentContent, setCurrentContent] = useState({ id: "info", move: Styles.moveInfo})
    const [user, setUser] = useState<UserInformation>()

    React.useEffect(() => {
        document.body.classList.toggle("no-scroll")

        const fetchData = async () => {
            try {
                const response = await axios.get<UserInformation>(`http://localhost:8000/user/${Username}`, {});
                setUser(response.data)
            }
            catch (error) {
                console.log(error)
            }   

        }

        fetchData();


        return () => {
            document.body.classList.remove("no-scroll")
        }


    },[]) 

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
                    <li>NedID: {user?.NetID ? user.NetID : "NA"}</li>
                    <li>Email: {user?.Email ? user.Email : "NA"}</li>
                    <li>URID: {user?.URID ? user.URID : "NA"}</li>
                    <li>Relationship Status: {user?.RelationshipStatus ? user.RelationshipStatus : "NA"}</li>
                    <li>Department: {user?.Department ? user.Department : "NA"}</li>
                    <li>Title: {user?.Title ? user.Title : "NA"}</li>
                    <li>Phone: {user?.Phone ? user.Phone : "NA"}</li>
                    <li>Location: {user?.Location ? user.Location : "NA"}</li>
                    <li>Password Set: {user?.LastPasswordSet ? user.LastPasswordSet : "NA"}</li>
                    <li>OU: {user?.OU ? user.OU : "NA"}</li>
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
import type React from "react"
import type { UserInformation } from "../models/User"



export const UserAbout: React.FC<{ user?: UserInformation}> = ({ user = null }) => {
    return (
    <>
        <ul id="info">
            <li>
                <span>Username: </span>
                <span>{user?.Username ? user.Username : "NA"}</span>
            </li>
            <li>
                <span>NetID:</span> 
                <span>{user?.NetID ? user.NetID : "NA"}</span>
            </li>
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
    </>
    )
}

export const UserGroups: React.FC<{ user?: UserInformation }> = ({ user = null }) => {
    return <>
        (Groups) To Be Setup
        {user?.Username}
        </>
}

export const UserShareDrive: React.FC<{ user?: UserInformation }> = ({ user = null }) => {
        return <>
        (Share Drive) To Be Setup
        {user?.Username}
        </>
}





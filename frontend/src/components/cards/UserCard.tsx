import Styles from "../../styles/components/UserCard.module.css";
import { type UserCardComponent } from "../../models/User";
import Button from "../../styles/global/Button.module.css"
import { useContextMenu } from "../ContextMenuProvider";
import { useState } from "react";
import Menu from "./Menu";
import { useNavigate } from "react-router-dom";

const UserCard: React.FC<UserCardComponent> = ({ Name = "NA", Username = "NA" , NetID = "NA", Email = "NA", OU = "NA", Select, ItemsSelected = []  }) => {
     const { openMenu } = useContextMenu(); 
     const [ copied, setCopied ] = useState("");
     const navigate = useNavigate();
     const MenuValue = new Map<string, string>([
        ["Name", Name],
        ["Username", Username],
        ["NetID", NetID],
        ["Email", Email],
        ["OU", OU],
     ])



    const handleRightClick = (e: React.MouseEvent) => {
        e.preventDefault();
        openMenu(e.clientX, e.clientY, (<Menu items={MenuValue} setCopied={setCopied}/>));
    };

    const navigateToUserPage = () => {
        navigate(`/user/${Username}`)
    }



    return (
            <div className={Styles.card} onContextMenu={handleRightClick} onClick={navigateToUserPage}>
                <p>Username: {Username} </p>
                <p>Name: {Name} </p>
                <p>NetID: {NetID} </p>
                <p>Email: {Email} </p>
                <p>OU: {OU} </p>
                <button 
                    className={`${Button.unselected} ${ItemsSelected.includes(Name) ? Button.selected : ""}`}    
                    onClick={() => Select("User", Name)}>
                </button>

                {copied != "" ? copied : ""}
            </div>
    )
}

  export default UserCard;
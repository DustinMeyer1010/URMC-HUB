import Styles from "../../styles/components/UserCard.module.css";
import { type UserCardComponent } from "../../models/User";
import Button from "../../styles/global/Button.module.css"
import { useContextMenu } from "../ContextMenuProvider";
import { useState } from "react";
import Menu from "./Menu";

const UserCard: React.FC<UserCardComponent> = ({ Name = "NA", Username = "NA" , NetID = "NA", Email = "NA", OU = "NA", Select, ItemsSelected = []  }) => {
     const { openMenu } = useContextMenu(); 
     const [ copied, setCopied ] = useState("");
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



    return (
            <div className={Styles.card} onContextMenu={handleRightClick}>
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
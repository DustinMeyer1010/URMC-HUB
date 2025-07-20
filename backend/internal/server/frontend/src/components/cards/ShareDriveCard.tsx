import type React from "react";
import type { ShareDriveComponent } from "../../models/ShareDrive";
import Style from "../../styles/components/ShareDriveCard.module.css"
import { useState } from "react";
import { useContextMenu } from "../ContextMenuProvider";
import Menu from "./Menu";
import Button from "../../styles/global/Button.module.css"


const ShareDriveCard: React.FC<ShareDriveComponent> = ({ Name = "NA", Server = "NA", Groups = [], Select, ItemsSelected}) => {

 const { openMenu } = useContextMenu();
    const [ copied, setCopied ] = useState("");
    const MenuValues = new Map<string, string>([
        ["Name", Name],
        ["Server", Server],
    ])
    
    const handleRightClick = (e: React.MouseEvent) => {
        e.preventDefault();
        openMenu(e.clientX, e.clientY, (<Menu items={MenuValues} setCopied={setCopied}/>));
    }


    return (
        <div className={Style.card} onContextMenu={handleRightClick}>
            <p>Drive: {"\\\\" + Server + "\\" + Name }</p>
            <button 
                className={`${Button.unselected} ${ItemsSelected.includes(Name) ? Button.selected : ""}`}    
                onClick={() => Select("Computer", Name)}>
            </button>
            {copied}
        </div>
    )
}

export default ShareDriveCard;
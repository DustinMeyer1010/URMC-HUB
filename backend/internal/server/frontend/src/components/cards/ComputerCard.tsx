import React, { useState } from "react";
import { type ComputerCardComponent } from "../../models/Computer";
import Styles from "../../styles/components/ComputerCard.module.css"
import Button from "../../styles/global/Button.module.css"
import { useContextMenu } from "../ContextMenuProvider";
import Menu from "./Menu";

const ComputerCard: React.FC<ComputerCardComponent> = ( { Name = "NA", OU = "NA", OperatingSystem = "NA", Select, ItemsSelected = [] }  ) => {
    const { openMenu } = useContextMenu();
    const [ copied, setCopied ] = useState("");
    const MenuValues = new Map<string, string>([
        ["Name", Name],
        ["OU", OU],
        ["OperatingSystem", OperatingSystem]
    ])


    const handleRightClick = (e: React.MouseEvent) => {
        e.preventDefault();
        openMenu(e.clientX, e.clientY, (<Menu items={MenuValues}  setCopied={setCopied}/>));
    };


    
    return (
        <div className={Styles.card} onContextMenu={handleRightClick} >
            <p>Name: { Name }</p>
            <p>Operating System: { OperatingSystem }</p>
            <p>OU: { OU }</p>
            <button 
                className={`${Button.unselected} ${ItemsSelected.includes(Name) ? Button.selected : ""}`}    
                onClick={() => Select("Computer", Name)}>
            </button>
            {copied}

        </div>
    ) 
}

export default ComputerCard;
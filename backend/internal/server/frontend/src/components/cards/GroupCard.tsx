import type React from "react"
import type { GroupCardComponent } from "../../models/Group"
import Styles from "../../styles/components/GroupCard.module.css"
import Button from "../../styles/global/Button.module.css"
import { useContextMenu } from "../ContextMenuProvider"
import { useState } from "react"
import Menu from "./Menu"



const GroupCard: React.FC<GroupCardComponent> = ( { Name = "NA", OU = "NA", Description = "NA", Information = "NA", Select, ItemsSelected = [] } ) => {

    const { openMenu } = useContextMenu();
    const [ copied, setCopied ] = useState("");
    const MenuValue = new Map<string, string>([
        ["Name", Name],
        ["OU", OU],
        ["Description", Description],
        ["Information", Information],
     ])


    const handleRightClick = (e: React.MouseEvent) => {
        e.preventDefault();
        openMenu(e.clientX, e.clientY, (<Menu items={MenuValue} setCopied={setCopied}/>));
    };


    return (
        <div className={Styles.card} onContextMenu={handleRightClick}>
            <p>Name: {Name} </p>
            <p>Description: {Description} </p>
            <p>Information: {Information} </p>
            <p>OU: {OU} </p>
            <button 
                className={`${Button.unselected} ${ItemsSelected.includes(Name) ? Button.selected : "" }`}    
                onClick={() => Select("Group", Name)}>
            </button>
            {copied}
        </div>

    )
}

export default GroupCard
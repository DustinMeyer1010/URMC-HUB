import type React from "react"
import type { GroupCardComponent } from "../../models/Group"
import Styles from "../../styles/components/GroupCard.module.css"
import Button from "../../styles/global/Button.module.css"
import { useContextMenu } from "../ContextMenuProvider"
import { useState } from "react"
import { HandleCopy } from "./Copy"
import Menu from "../../styles/global/Menu.module.css"


const GroupCard: React.FC<GroupCardComponent> = ( { Name = "NA", OU = "NA", Description = "NA", Information = "NA", Select, ItemsSelected = [] } ) => {

    const { openMenu } = useContextMenu();
    const [ copied, setCopied ] = useState("");


    const handleRightClick = (e: React.MouseEvent) => {
    e.preventDefault();
    openMenu(e.clientX, e.clientY, (
      <div className={Menu.default}>
        <h1>Copy Options</h1>
        <ul style={{ margin: 0, padding: 0, listStyle: "none" }}>
            <li >
                <button onClick={() => HandleCopy("Name", Name, setCopied)}>Name</button>
            </li>
            <li>
              <button onClick={() => HandleCopy("OU", OU, setCopied)}>OS</button>
            </li>
            <li >
              <button onClick={() => HandleCopy("Description", Description, setCopied)}>Description</button>
            </li>
            <li >
              <button onClick={() => HandleCopy("Information", Information, setCopied)}>information</button>
            </li>
        </ul>
      </div>
    ));
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
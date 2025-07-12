import React, { useState } from "react";
import { CopyComputerAtrribute, type ComputerCardComponent } from "../../models/Computer";
import Styles from "../../styles/components/ComputerCard.module.css"
import Button from "../../styles/global/Button.module.css"
import ComputerMenuStyles from "../../styles/components/ComputerMenu.module.css"
import { useContextMenu } from "../ContextMenuProvider";

const ComputerCard: React.FC<ComputerCardComponent> = ( { Name = "NA", OU = "NA", OperatingSystem = "NA", Select, ItemsSelected = [] }  ) => {
    const { openMenu } = useContextMenu();
    const [ copied, setCopied ] = useState("");

    const HandleCopy = (copyItem: string) => {

      if (copied != "") {
        return
      }

      switch (copyItem) {
        case "name": {
          CopyComputerAtrribute(Name)
          setCopied("Name Copied")
          break

        }
        case "os": {
          CopyComputerAtrribute(OperatingSystem)
          setCopied("OS Copied")
          break
        }
        case "ou": {
          CopyComputerAtrribute(OU)
          setCopied("OU Copied")
          break
        }
        case "all": {
          const text = `Name: ${Name}\nOU: ${OU}\nOS: ${OperatingSystem}`
          CopyComputerAtrribute(text)
          setCopied("All Copied")
          break
        }
      }
        setTimeout(() => {
            setCopied("")
          }, 3000)
    }



    const handleRightClick = (e: React.MouseEvent) => {
    e.preventDefault();
    openMenu(e.clientX, e.clientY, (
      <div className={ComputerMenuStyles.default}>
        <h1>Copy Options</h1>
        <ul style={{ margin: 0, padding: 0, listStyle: "none" }}>
            <li >
                <button onClick={() => HandleCopy("name")}>Name</button>
            </li>
            <li>
              <button onClick={() => HandleCopy("os")}>OS</button>
            </li>
            <li >
              <button onClick={() => HandleCopy("ou")}>OU</button>
            </li>
            <li >
              <button onClick={() => HandleCopy("all")}>All</button>
            </li>
        </ul>
      </div>
    ));
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
            {copied != "" ? copied : ""}

        </div>
    ) 
}

export default ComputerCard;
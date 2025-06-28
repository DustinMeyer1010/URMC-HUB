import React from "react";
import type { ComputerCardComponent } from "../../models/Computer";
import Styles from "../../styles/components/ComputerCard.module.css"
import Button from "../../styles/global/Button.module.css"

const ComputerCard: React.FC<ComputerCardComponent> = ( { Name = "NA", OU = "NA", OperatingSystem = "NA", Select, ItemsSelected = [] }  ) => {

    
    return (
        <div className={Styles.card}>
            <p>Name: { Name }</p>
            <p>Operating System: { OperatingSystem }</p>
            <p>OU: { OU }</p>
            <button 
                className={`${Button.unselected} ${ItemsSelected.includes(Name) ? Button.selected : ""}`}    
                onClick={() => Select("Computer", Name)}>
            </button>
        </div>
    ) 
}

export default ComputerCard;
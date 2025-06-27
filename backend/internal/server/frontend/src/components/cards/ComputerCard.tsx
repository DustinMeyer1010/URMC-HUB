import React from "react";
import type { ComputerCardComponent } from "../../models/Computer";
import ComputerCardStyles from "../../styles/components/ComputerCard.module.css"

const ComputerCard: React.FC<ComputerCardComponent> = ( { Name = "NA", OU = "NA", OperatingSystem = "NA", Select, ItemsSelected = [] }  ) => {

    
    return (
        <div className={ComputerCardStyles.card}>
            <p>Name: { Name }</p>
            <p>Operating System: { OperatingSystem }</p>
            <p>OU: { OU }</p>
            <button 
                className={`${ComputerCardStyles.button} ${ItemsSelected.includes(Name) ? ComputerCardStyles.button_selected : ComputerCardStyles.button_unselected}`}    
                onClick={() => Select("Computer", Name)}></button>
        </div>
    ) 
}

export default ComputerCard;
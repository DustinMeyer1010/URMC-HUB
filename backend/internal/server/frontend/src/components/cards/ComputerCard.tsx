import React from "react";
import type { ComputerCardComponent } from "../../models/Computer";
import ComputerCardStyles from "../../styles/components/ComputerCard.module.css"

const ComputerCard: React.FC<ComputerCardComponent> = ( { Name = "NA", OU = "NA", OperatingSystem = "NA", Select }  ) => {
    
    return (
        <div className={ComputerCardStyles.card}>
            <p>Name: { Name }</p>
            <p>Operating System: { OperatingSystem }</p>
            <p>OU: { OU }</p>
            <button onClick={() => Select("Computer", Name)}></button>
        </div>
    ) 
}

export default ComputerCard;
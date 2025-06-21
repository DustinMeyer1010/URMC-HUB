import React from "react";
import type { ComputerCardInfo } from "../../models/Computer";
import ComputerCardStyles from "../../styles/components/ComputerCard.module.css"

const ComputerCard: React.FC<ComputerCardInfo> = ( { Name = "NA", OU = "NA", OperatingSystem = "NA" } ) => {
    
    return (
        <div className={ComputerCardStyles.card}>
            <p>Name: { Name }</p>
            <p>Operating System: { OperatingSystem }</p>
            <p>OU: { OU }</p>
        </div>
    ) 
}

export default ComputerCard;
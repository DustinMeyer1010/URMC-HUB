import React from "react";
import type { ComputerCardInfo } from "../../models/Computer";
import ComputerCardStyles from "../../styles/components/ComputerCard.module.css"

const ComputerCard: React.FC<ComputerCardInfo> = ( { Name = "", OU = "" } ) => {
    
    return (
        <div className={ComputerCardStyles.card}>
            <p>Name: { Name }</p>
            <p>OU: { OU }</p>
        </div>
    ) 
}

export default ComputerCard;
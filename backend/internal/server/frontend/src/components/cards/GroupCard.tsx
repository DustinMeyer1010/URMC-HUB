import type React from "react"
import type { GroupCardInfo } from "../../models/Group"
import GroupCardStyle from "../../styles/components/GroupCard.module.css"


const GroupCard: React.FC<GroupCardInfo> = ( { Name = "NA", OU = "NA", Description = "NA", Information = "NA" } ) => {

    return (
        <div className={GroupCardStyle.card}>
        <p>Name: {Name} </p>
        <p>Description: {Description} </p>
        <p>Information: {Information} </p>
        <p>OU: {OU} </p>
        </div>
    )
}

export default GroupCard
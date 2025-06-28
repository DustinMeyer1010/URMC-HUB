import type React from "react"
import type { GroupCardComponent } from "../../models/Group"
import Styles from "../../styles/components/GroupCard.module.css"
import Button from "../../styles/global/Button.module.css"


const GroupCard: React.FC<GroupCardComponent> = ( { Name = "NA", OU = "NA", Description = "NA", Information = "NA", Select, ItemsSelected = [] } ) => {

    return (
        <div className={Styles.card}>
        <p>Name: {Name} </p>
        <p>Description: {Description} </p>
        <p>Information: {Information} </p>
        <p>OU: {OU} </p>
        <button 
                className={`${Button.unselected} ${ItemsSelected.includes(Name) ? Button.selected : "" }`}    
                onClick={() => Select("Group", Name)}>
        </button>
        </div>
    )
}

export default GroupCard
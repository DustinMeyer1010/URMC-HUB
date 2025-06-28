import Styles from "../../styles/components/UserCard.module.css";
import type { UserCardComponent } from "../../models/User";
import Button from "../../styles/global/Button.module.css"

const UserCard: React.FC<UserCardComponent> = ({ Name = "NA", Username = "NA" , NetID = "NA", Email = "NA", OU = "NA", Select, ItemsSelected = []  }) => {

    return (
            <div className={Styles.card}>
                <p>Username: {Username} </p>
                <p>Name: {Name} </p>
                <p>NetID: {NetID} </p>
                <p>Email: {Email} </p>
                <p>OU: {OU} </p>
                <button 
                    className={`${Button.unselected} ${ItemsSelected.includes(Name) ? Button.selected : ""}`}    
                    onClick={() => Select("User", Name)}>
                </button>
            </div>
    )
}

  export default UserCard;
import UserCardStyles from "../../styles/components/UserCard.module.css";
import type { UserCardInfo } from "../../models/User";

const UserCard: React.FC<UserCardInfo> = ({ Name = "NA", Username = "NA" , NetID = "NA", Email = "NA", OU = "NA" }) => {

    return (
            <div className={UserCardStyles.card}>
                <p>Username: {Username} </p>
                <p>Name: {Name} </p>
                <p>NetID: {NetID} </p>
                <p>Email: {Email} </p>
                <p>OU: {OU} </p>
            </div>
    )
}

  export default UserCard;
import "./UserCard.css";

type UserCardInfo = {
    Name?: String,
    Username?: String,
    NetID?: String,
    Email?: String,
    OU?: String,
}

const UserCard: React.FC<UserCardInfo> = ({ Name = "", Username = "" , NetID = "", Email = "", OU = "" }) => {
    

    return (
            <div className="user-container">
                <p>Username: {Username} </p>
                <p>Name: {Name} </p>
                <p>NetID: {NetID} </p>
                <p>Email: {Email} </p>
                <p>OU: {OU} </p>
            </div>
    )
}

  export default UserCard;
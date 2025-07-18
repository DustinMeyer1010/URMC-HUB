import Styles from "../../styles/components/UserCard.module.css";
import { CopyUserAtrribute, type UserCardComponent } from "../../models/User";
import Button from "../../styles/global/Button.module.css"
import { useContextMenu } from "../ContextMenuProvider";
import UserMenuStyle from "../../styles/components/UserMenu.module.css"
import { useState } from "react";

const UserCard: React.FC<UserCardComponent> = ({ Name = "NA", Username = "NA" , NetID = "NA", Email = "NA", OU = "NA", Select, ItemsSelected = []  }) => {
     const { openMenu } = useContextMenu(); 
     const [ copied, setCopied ] = useState("");


 const HandleCopy = (copyItem: string) => {

      if (copied != "") {
        return
      }

      switch (copyItem) {
        case "name":
            CopyUserAtrribute(Name)
            setCopied("Name Copied")
            break
        case "netid":
            CopyUserAtrribute(NetID)
            setCopied("NetID Copied")
            break
        case "email":
            CopyUserAtrribute(Email)
            setCopied("Email Copied")
            break
        case "ou":
            CopyUserAtrribute(OU)
            setCopied("OU Copied")
            break
        case "username":
            CopyUserAtrribute(Username)
            setCopied("All Copied")
            break
        case "all":
            CopyUserAtrribute(Name)
            setCopied("All Copied")
            break
      }
        setTimeout(() => {
            setCopied("")
          }, 3000)
    }



    const handleRightClick = (e: React.MouseEvent) => {
    e.preventDefault();
    openMenu(e.clientX, e.clientY, (
      <div className={UserMenuStyle.default}>
        <h1>Copy Options</h1>
        <ul style={{ margin: 0, padding: 0, listStyle: "none" }}>
            <li >
                <button onClick={() => HandleCopy("username")}>Username</button>
            </li>
            <li >
                <button onClick={() => HandleCopy("name")}>Name</button>
            </li>
            <li>
              <button onClick={() => HandleCopy("netid")}>NetID</button>
            </li>
            <li >
              <button onClick={() => HandleCopy("email")}>Email</button>
            </li>
            <li >
              <button onClick={() => HandleCopy("ou")}>OU</button>
            </li>
            <li >
              <button onClick={() => HandleCopy("all")}>All</button>
            </li>
        </ul>
      </div>
    ));
  };



    return (
            <div className={Styles.card} onContextMenu={handleRightClick}>
                <p>Username: {Username} </p>
                <p>Name: {Name} </p>
                <p>NetID: {NetID} </p>
                <p>Email: {Email} </p>
                <p>OU: {OU} </p>
                <button 
                    className={`${Button.unselected} ${ItemsSelected.includes(Name) ? Button.selected : ""}`}    
                    onClick={() => Select("User", Name)}>
                </button>

                {copied != "" ? copied : ""}
            </div>
    )
}

  export default UserCard;
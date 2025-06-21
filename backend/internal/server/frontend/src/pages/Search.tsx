
import { useState } from "react";

//import UserCard from "../components/cards/UserCard";
import SearchBox from "../components/SearchBox";
import Paging from "../components/Paging";
//import ComputerCard from "../components/cards/ComputerCard"

import { UserSearch, GroupSearch } from "./SearchHelper"

import type { UserCardInfo } from "../models/User";

import SearchStyles from "../styles/pages/Search.module.css"
import HomeStyles from "../styles/pages/Home.module.css";
import GroupCard from "../components/cards/GroupCard";
import type { GroupCardInfo } from "../models/Group";

function Search() {
    const [searchValue, setSearchValue] = useState("");
    const [usersResults, setUsersResults] = useState<UserCardInfo[]>([])
    const [groupsResults, setGroupsResults] = useState<GroupCardInfo[]>([])
    const [currentPage, setCurrentPage] = useState(1)
    const itemsPerPage = 5;
    const idxLastItem = currentPage * itemsPerPage
    const idxFirstItem = idxLastItem - itemsPerPage

    // Changed based 
    //const currentUsersItems = usersResults.slice(idxFirstItem, idxLastItem)
    const currentGroupItems = groupsResults.slice(idxFirstItem, idxLastItem)

    const totalPage = Math.ceil(groupsResults.length + usersResults.length / itemsPerPage)

    const handleChange = (newValue: string) => {
        setSearchValue(newValue);
    }

     const handleEnter = async (e: React.KeyboardEvent<HTMLInputElement>) => {
        if (e.key != 'Enter') {
            console.log("Not Enter");
            return;
        }
        setCurrentPage(1)
        handleSubmit();

        };
    
    const handleSubmit = async () => {
        setGroupsResults(await GroupSearch(searchValue));
        setUsersResults(await UserSearch(searchValue));
    };



    return (
        <div>
            <div className={SearchStyles.results_container}>
                {
                    /*
                currentUsersItems.map((user) => (
                    <UserCard Name={user.Name} Username={user.Username} NetID={user.NetID} Email={user.Email} OU={user.OU}/>
                ))
                    */
                }
                {
                currentGroupItems.map((group) => (
                    <GroupCard Name={group.Name} OU={group.OU} Description={group.Description} Information={group.Information} />
                ))
                }
            </div>
            <div className={HomeStyles.control_container}>
                <SearchBox
                    value={searchValue}
                    onChange={handleChange}
                    onKeyDown={handleEnter}
                    onClick={handleSubmit}
                    />
                
                <Paging currentPage={currentPage} totalPages={totalPage} onPageChange={setCurrentPage}/>
            </div>

        </div>
    )
}

  export default Search;
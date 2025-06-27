
import { useState } from "react";

//import UserCard from "../components/cards/UserCard";
import SearchBox from "../components/SearchBox";
import Paging from "../components/Paging";

//import ComputerCard from "../components/cards/ComputerCard"
//import GroupCard from "../components/cards/GroupCard";
import type { ComputerCardInfo } from "../models/Computer";

import { UserSearch, GroupSearch, ComputerSearch } from "./SearchHelper"

import type { UserCardInfo } from "../models/User";

import SearchStyles from "../styles/pages/Search.module.css"
import HomeStyles from "../styles/pages/Home.module.css";
import type { GroupCardInfo } from "../models/Group";
import ComputerCard from "../components/cards/ComputerCard";
import SideBar from "../components/SideBar";

function Search() {
    const [searchValue, setSearchValue] = useState("");
    const [usersResults, setUsersResults] = useState<UserCardInfo[]>([])
    const [groupsResults, setGroupsResults] = useState<GroupCardInfo[]>([])
    const [computerResults, setComputerResults] = useState<ComputerCardInfo[]>([])
    const [selected, setSelected] = useState<string[]>([])
    const [currentPage, setCurrentPage] = useState(1)
    const itemsPerPage = 5;
    const idxLastItem = currentPage * itemsPerPage
    const idxFirstItem = idxLastItem - itemsPerPage

    // Changed based 
    //const currentUsersItems = usersResults.slice(idxFirstItem, idxLastItem)
    //const currentGroupItems = groupsResults.slice(idxFirstItem, idxLastItem)
    const currentCompterItems = computerResults.slice(idxFirstItem, idxLastItem)

    const totalPage = Math.ceil(groupsResults.length + usersResults.length + computerResults.length / itemsPerPage)

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
        setComputerResults(await ComputerSearch(searchValue))
    };

    const handleSelection = (Type: string, Name: string) => {
        setSelected((item) => item.includes(Name) ? item.filter((i) => i !== Name) : [...item, Name]);
    };

    const handleClear = () => {
        setSelected([])
    }



    return (
        <div>
            <div className={SearchStyles.results_container}>
                <SideBar clear={handleClear} selected={selected} setSelected={handleSelection}/>
                <ComputerCard Select={handleSelection} ItemsSelected={selected}/>
                
                <ComputerCard Name="Test1" Select={handleSelection} ItemsSelected={selected}/>
                <ComputerCard Name="Test2" Select={handleSelection} ItemsSelected={selected}/>
                <ComputerCard Name="Test3" Select={handleSelection} ItemsSelected={selected}/>
                <ComputerCard Name="Test5" Select={handleSelection} ItemsSelected={selected}/>
                {
                    /*
                currentUsersItems.map((user) => (
                    <UserCard Name={user.Name} Username={user.Username} NetID={user.NetID} Email={user.Email} OU={user.OU}/>
                ))
                    */
                }
                {
                    /*
                currentGroupItems.map((group) => (
                    <GroupCard Name={group.Name} OU={group.OU} Description={group.Description} Information={group.Information} />
                ))
                    */
                }
                {
                currentCompterItems.map((computer) => (
                    <ComputerCard Name={computer.Name} OU={computer.OU} OperatingSystem={computer.OperatingSystem} Select={handleSelection}  ItemsSelected={selected}/>
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
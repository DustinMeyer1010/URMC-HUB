
import { useState } from "react";

import UserCard from "../components/cards/UserCard";
import SearchBox from "../components/SearchBox";
import Paging from "../components/Paging";
import ComputerCard from "../components/cards/ComputerCard"

import { UserSearch } from "./SearchHelper"

import type { UserCardInfo } from "../models/User";

import SearchStyles from "../styles/pages/Search.module.css"
import HomeStyles from "../styles/pages/Home.module.css";

function Search() {
    const [searchValue, setSearchValue] = useState("");

    const handleChange = (newValue: string) => {
        setSearchValue(newValue);
    }

     const handleEnter = async (e: React.KeyboardEvent<HTMLInputElement>) => {
        if (e.key != 'Enter') {
            console.log("Not Enter");
            return;
        }
        handleSubmit();

        };
    
    const handleSubmit = async () => {
        let test: UserCardInfo[] = await UserSearch(searchValue);

        console.log(test)
        
    };

    return (
        <div>
            <div className={SearchStyles.results_container}>
                <UserCard/>
                <ComputerCard/>
                <UserCard/>
                <UserCard/>
            </div>
            <div className={HomeStyles.control_container}>
                <SearchBox
                    value={searchValue}
                    onChange={handleChange}
                    onKeyDown={handleEnter}
                    onClick={handleSubmit}
                    />
                
                <Paging/>
            </div>

        </div>
    )
}

  export default Search;
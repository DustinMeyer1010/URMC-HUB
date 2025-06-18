
import { useState } from "react";
import axios from "axios";
import UserCard from "../components/UserCard";
import SearchBox from "../components/SearchBox";
import HomeStyles from "./Home.module.css";
import Paging from "../components/Paging";
import SearchStyles from "./Search.module.css"

function Search() {
    const [searchValue, setSearchValue] = useState("");
    const [data, setData] = useState<any>(null);

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
        
        try {
            const response = await axios.get(`http://localhost:8080/search/users/${searchValue}`, {});
                setData(response.data); 
            
            
          } catch (error) {
            if (axios.isAxiosError(error)){
                if (error.response) {
                    console.log(error.response.data)
                }
            } else {
                console.error(error)
            }
          }
    };

    return (
        <div>
            <div className={SearchStyles.results_container}>
                <UserCard/>
                
                <UserCard/>
                <UserCard/>
                <UserCard/>
                <UserCard/>
                <UserCard/>
                <UserCard/>
                <UserCard/>
                <UserCard/>
                <UserCard/>
                <UserCard/>
                <UserCard/>
                <UserCard/>
                {JSON.stringify(data)}
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

import { useState } from "react";
import axios from "axios";
import UserCard from "../components/UserCard";
import SearchBox from "../components/SearchInput";
import HomeStyles from "./Home.module.css";
import Button from "../styles/Button.module.css";
import SearchStyles from "./Search.module.css";
import Paging from "../components/Paging";

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
            console.error('Error fetching data:', error);
          }
    };

    return (
        <div>
            <div className="results-container">
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
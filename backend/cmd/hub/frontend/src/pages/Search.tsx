
import { useState } from "react";
import "./Search.css"
import axios from "axios";
import UserCard from "../components/UserCard";
import SearchBox from "../components/SearchBox";

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
            <div className="control-container">
            <button onClick={handleSubmit}>Search</button>
            <SearchBox
                value={searchValue}
                onChange={handleChange}
                onKeyDown={handleEnter}
                />
           
                <button className="pager-button">-</button>
                <button className="pager-button">+</button>

            </div>

        </div>
    )
}

  export default Search;
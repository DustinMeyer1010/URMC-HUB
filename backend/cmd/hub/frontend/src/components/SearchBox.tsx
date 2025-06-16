import SearchBoxStyles from "./SearchBox.module.css"


type SearchBox = {
    value: string;
    onChange: (newValue: string) => void;
    onKeyDown: (e: React.KeyboardEvent<HTMLInputElement>) => void;
}

const SearchBox: React.FC<SearchBox> = ({value, onChange, onKeyDown}) => {
    return (
        <div>
            <input 
                className={SearchBoxStyles.input} 
                type="text" 
                value={value}
                placeholder="Search"
                onChange={(e) => onChange(e.target.value)} 
                onKeyDown={onKeyDown}
                />
        </div>
    )
}


export default SearchBox;
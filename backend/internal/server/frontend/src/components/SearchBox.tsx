import Input from "../styles/global/Input.module.css"
import Button from "../styles/global/Button.module.css"


type SearchBox = {
    value: string;
    onChange: (newValue: string) => void;
    onKeyDown: (e: React.KeyboardEvent<HTMLInputElement>) => void;
    onClick: () => void;
}

const SearchBox: React.FC<SearchBox> = ({value, onChange, onKeyDown, onClick}) => {
    return (
        <div style={{display: 'flex', gap: '10px'}}>
            <button className={Button.default} onClick={onClick}>Search</button>
            <input 
                className={Input.default} 
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
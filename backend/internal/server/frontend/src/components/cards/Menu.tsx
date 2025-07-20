import type React from "react"
import { HandleCopy } from "./Copy"
import Style from "../../styles/global/Menu.module.css"

type MenuModel = {
    items: Map<string, string>
    setCopied: React.Dispatch<React.SetStateAction<string>>
}

const Menu: React.FC<MenuModel> = ({ items,  setCopied }) => {


    return (
        <div className={Style.default}>
            <h1>Copy Options</h1>
            <ul style={{ margin: 0, padding: 0, listStyle: "none" }}>
            {
                [...items.entries()].map(([key, value]) => (
                <li >
                    <button onClick={() => HandleCopy(key, value, setCopied)}>{key}</button>
                </li>
                ))
            }   
            </ul>
        </div>
    )

}

export default Menu;
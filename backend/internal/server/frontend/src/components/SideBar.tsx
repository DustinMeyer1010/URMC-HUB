
import type React from "react"
import styles from "../styles/components/SideBar.module.css"
import Button from "../styles/global/Button.module.css"
import { useState } from "react"

type Selection = {
    selected: string[]
    setSelected: (Type: string, Name: string) => void
    clear: () => void
}

const SideBar: React.FC<Selection> = ({selected, setSelected, clear}) => {
    const [hidden, setHidden] = useState(370)

    const toggleMenu = () => {
        if (hidden == 370){
            setHidden(-20)
            return
        } 
        setHidden(370)

    }

    return (
        <div style={{transform: `translateY(${hidden}px)`}} className={styles.container}>
            <button onClick={toggleMenu} className={`${styles.tab_button} ${Button.default}`}>Items Menu</button>
            <div className={styles.list_container}>
                {selected.map((items) =>
                <div className={styles.item}>{items} <button className={styles.item_button} onClick={() => setSelected("computer", items)}>X</button></div>
                )}
            </div>
            <div className={styles.button_container}> 
                    <button className={`${Button.default} ${styles.search_button}`}>Search All</button>
                <button onClick={clear} className={`${Button.default} ${styles.clear_button}`}>Clear All</button>
            </div>
        </div>
    )
}

export default SideBar
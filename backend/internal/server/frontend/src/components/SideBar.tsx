
import type React from "react"
import styles from "../styles/components/SideBar.module.css"

type Selection = {
    selected: string[]
    setSelected: (Type: string, Name: string) => void
}

const SideBar: React.FC<Selection> = ({selected, setSelected}) => {

    return (
        <div className={styles.container}>
            {selected.map((items) =>
            <div>{items} <button onClick={() => setSelected("computer", items)}></button></div>
            )}
        </div>
    )
}

export default SideBar
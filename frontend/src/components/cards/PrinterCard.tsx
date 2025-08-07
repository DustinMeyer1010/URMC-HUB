import type React from "react";
import type { PrinterCardComponent } from "../../models/Printer";
import Style from "../../styles/components/PrinterCard.module.css"
import { useContextMenu } from "../ContextMenuProvider";
import { useState } from "react";
import Button from "../../styles/global/Button.module.css"
import Menu from "./Menu";


const PrinterCard: React.FC<PrinterCardComponent> = ({ Name = "NA", Model = "NA", IP = "000.000.000.000", Processor = "NA", Location = "NA", Notes = "NA", Select ,ItemsSelected = []}) => {
    const { openMenu } = useContextMenu();
    const [ copied, setCopied ] = useState("");
    const MenuValues = new Map<string, string>([
        ["Name", Name],
        ["Model", Model],
        ["IP", IP],
        ["Processor", Processor],
        ["Location", Location],
        ["Notes", Notes]
    ])


    
    const handleRightClick = (e: React.MouseEvent) => {
        e.preventDefault();
        openMenu(e.clientX, e.clientY, (<Menu items={MenuValues} setCopied={setCopied}/>));
    }


    return (
        <div className={Style.card} onContextMenu={handleRightClick}>
            <p>{ Name }</p>
            <p>{ Model }</p>
            <p>{ IP }</p>
            <p>{ Processor }</p>
            <p>{ Location }</p>
            <p>{ Notes }</p>

            <button 
                className={`${Button.unselected} ${ItemsSelected.includes(Name) ? Button.selected : ""}`}    
                onClick={() => Select("Computer", Name)}>
            </button>
            {copied}
        </div>
    )
}


export default PrinterCard;
import type React from "react";
import type { Printer } from "../../models/Printer";
import Style from "../../styles/components/PrinterCard.module.css"

const PrinterCard: React.FC<Printer> = ({ Name = "NA", Model = "NA", IP = "000.000.000.000", Location = "NA", Notes = "NA" }) => {

    return (
        <div className={Style.card}>
            <p>{ Name }</p>
            <p>{ Model }</p>
            <p>{ IP }</p>
            <p>{ Location }</p>
            <p>{ Notes }</p>
        </div>
    )
}

export default PrinterCard;
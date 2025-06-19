import React from "react";

import Button from "../styles/global/Button.module.css";

import PagingStyle from "../styles/components/Paging.module.css";


type PagingOptions = {
    currentPage: number;
    totalPages: number;
    onPageChange: (page: number) => void
}

const Paging: React.FC<PagingOptions> = ( { currentPage, totalPages, onPageChange}) => {
    return (
        <div className={PagingStyle.pager_container}>
            <span className={PagingStyle.page_numbers}>{currentPage} / {totalPages}</span>
            <div className={PagingStyle.pager_button_container}>
                <button onClick={() => {if(currentPage > 1) onPageChange(currentPage - 1)}} className={`${Button.default} ${PagingStyle.pager_button}`}>-</button>
                <button onClick={() => {if(currentPage < totalPages) onPageChange(currentPage + 1)}} className={`${Button.default} ${PagingStyle.pager_button}`}>+</button>
            </div>
        </div>
    )
}

export default Paging;
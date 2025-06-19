import Button from "../styles/global/Button.module.css";

import PagingStyle from "../styles/components/Paging.module.css";

const Paging = () => {
    return (
        <div className={PagingStyle.pager_container}>
            <span className={PagingStyle.page_numbers}>1 / 100</span>
            <div className={PagingStyle.pager_button_container}>
                <button className={`${Button.default} ${PagingStyle.pager_button}`}>-</button>
                <button className={`${Button.default} ${PagingStyle.pager_button}`}>+</button>
            </div>
        </div>
    )
}

export default Paging;
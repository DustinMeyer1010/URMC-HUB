
import { Link } from "react-router-dom"
import { useEffect, useState } from "react"

import NavStyles from "../styles/components/Nav.module.css"

const Nav = () => {

    const [scrolled, setScrolled] = useState(false);

    useEffect(() => {
        const onScroll = () => {
            setScrolled(window.scrollY > 1)
        }

        window.addEventListener('scroll', onScroll);
        return () => window.removeEventListener('scroll', onScroll)
    });





    return (
        <nav className={!scrolled ? NavStyles.nav : `${NavStyles.nav_top}`}>
          <Link className={NavStyles.a} to="/">Home</Link>
          <Link className={NavStyles.a} to="/search">Search</Link>
          <Link className={NavStyles.a} to="/SummaryIndex">System Summary Index</Link>

        </nav>
    )
}

export default Nav;
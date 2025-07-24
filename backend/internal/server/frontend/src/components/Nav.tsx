
import { Link } from "react-router-dom"
import { useEffect, useState } from "react"

import NavStyles from "../styles/components/Nav.module.css"

const Nav = () => {





    return (
        <nav className={`${NavStyles.nav_top}`}>
          <Link className={NavStyles.a} to="/">Home</Link>
          <Link className={NavStyles.a} to="/search">Search</Link>

        </nav>
    )
}

export default Nav;
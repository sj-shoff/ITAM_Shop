import { Link } from "react-router-dom"
import { linksList } from "../config/links-config"
import classes from "./navbar.module.scss"

export function Navbar() {
    return (
        <nav className={classes.nav}>
            <ul className={classes.linksList}>
                {linksList.map((el) => (
                    <Link key={el.path} to={el.path} className={classes.link}>
                        {el.value}
                    </Link>
                ))}
            </ul>
        </nav>
    )
}

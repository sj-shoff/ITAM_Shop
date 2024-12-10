import { Link } from "react-router-dom"
import classes from "./navbar.module.scss"
import { LinksConfiguration } from "../model/navbar-model"

type NavbarProps = {
    linksList: LinksConfiguration
}

export function Navbar({ linksList }: NavbarProps) {
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

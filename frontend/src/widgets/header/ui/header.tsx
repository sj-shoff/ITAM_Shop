import { Navbar } from "@widgets/navbar"
import { ButtonsGroup } from "@widgets/buttons-group"
import { Link } from "react-router-dom"
import { Logo } from "@shared/ui/logo"
import classes from "./header.module.scss"

export function Header() {
    return (
        <header className={classes.header}>
            <Navbar />
            <Link to='/'>
                <Logo />
            </Link>
            <ButtonsGroup />
        </header>
    )
}

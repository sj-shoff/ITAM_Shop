import { Main } from "@widgets/main"
import { Header } from "@widgets/header"
import classes from "./dafault-layout.module.scss"
import { linksList } from "../config/links-config"
import { Outlet } from "react-router-dom"

export function DefaultLayout() {
    return (
        <>
            <div className={classes.background} />
            <div className={classes.page}>
                <Header linksList={linksList} />
                <Main>
                    <Outlet />
                </Main>
            </div>
        </>
    )
}

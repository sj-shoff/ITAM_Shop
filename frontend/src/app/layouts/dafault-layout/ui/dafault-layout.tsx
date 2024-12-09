import { Main } from "@widgets/main"
import { ReactNode } from "react"
import { Header } from "@widgets/header"
import classes from "./dafault-layout.module.scss"
import { linksList } from "../config/links-config"

type defaultLayoutProps = {
    children: ReactNode
}

export function DefaultLayout({ children }: defaultLayoutProps) {
    return (
        <>
            <div className={classes.background} />
            <div className={classes.page}>
                <Header linksList={linksList} />
                <Main>{children}</Main>
            </div>
        </>
    )
}

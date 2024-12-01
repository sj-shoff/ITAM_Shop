import { Main } from "@widgets/main"
import { ReactNode } from "react"
import { Header } from "@widgets/header"
import classes from "./dafault-layout.module.scss"

type defaultLayoutProps = {
    children: ReactNode
}

export function DefaultLayout({ children }: defaultLayoutProps) {
    return (
        <>
            <div className={classes.background} />
            <div className={classes.page}>
                <Header />
                <Main>{children}</Main>
            </div>
        </>
    )
}

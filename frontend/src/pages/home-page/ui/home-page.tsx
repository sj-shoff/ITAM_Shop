import classes from "./home-page.module.scss"
import { ItemsList } from "@widgets/items-list"

export function Homepage() {
    return (
        <>
            <div className={classes.homepageBg}></div>
            <section>
                <ItemsList />
            </section>
        </>
    )
}

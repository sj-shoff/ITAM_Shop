import { ItemsList } from "@widgets/items-list"
import classes from "./home-page.module.scss"

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

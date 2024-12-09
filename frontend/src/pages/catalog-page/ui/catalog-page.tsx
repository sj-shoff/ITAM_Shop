import classes from "./home-page.module.scss"
import { ItemsList } from "@widgets/items-list"

export function CatalogPage() {
    return (
        <>
            <div className={classes.homepageBg}></div>
            <section>
                <ItemsList />
            </section>
        </>
    )
}

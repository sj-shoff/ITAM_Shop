import { Button } from "@nextui-org/react"
import classes from "./catalog-page.module.scss"
import { ItemsList } from "@widgets/items-list"
import { FilterOptions } from "@entities/product"

export function CatalogPage() {

    function clickHandler() {
        const opt: FilterOptions = {
            category: "Категория 4",
            minPrice: 0,
            maxPrice: 10000000,
        }

        console.log(opt)
    }

    return (
        <>
            <div className={classes.homepageBg}></div>
            <section>
                <Button onClick={clickHandler}>Filter</Button>
                <ItemsList />
            </section>
        </>
    )
}

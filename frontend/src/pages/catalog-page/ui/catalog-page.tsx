import { Button } from "@nextui-org/react"
import classes from "./catalog-page.module.scss"
import { ItemsList } from "@widgets/items-list"
import { FilterOptions, useFilterProductsMutation } from "@entities/product"

export function CatalogPage() {
    const [filterProducts] = useFilterProductsMutation()

    function clickHandler() {
        const opt: FilterOptions = {
            category: "Категория 4",
            minPrice: 0,
            maxPrice: 10000000,
        }

        filterProducts(opt)
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

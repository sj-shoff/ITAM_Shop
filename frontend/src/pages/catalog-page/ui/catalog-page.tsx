import { Button } from "@nextui-org/react"
import classes from "./catalog-page.module.scss"
import { ItemsList } from "@widgets/items-list"
import { FilterOptions, getList, useGetProductsQuery } from "@entities/product"
import { useState } from "react"
import { useSelector } from "react-redux"

export function CatalogPage() {
    // При применении фильтра компонент будет перерисовываться
    const [isFiltered, setIsFiltered] = useState<boolean>(false)
    const { data } = useGetProductsQuery()
    const filteredList = useSelector(getList)

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
            <section className={classes.content}>
                <Button onClick={clickHandler}>Filter</Button>
                <ItemsList data={isFiltered ? filteredList : data} />
            </section>
        </>
    )
}

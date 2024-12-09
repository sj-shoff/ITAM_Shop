import classes from "./items-list.module.scss"
import { productsMaper } from "../lib/maper"
import { useGetProductsQuery } from "@entities/product"
import { Spinner } from "@nextui-org/react"

export function ItemsList() {
    const { data } = useGetProductsQuery()

    return (
        <div className={classes.list}>
            {data ? productsMaper(data) : <Spinner label='loading...' />}
        </div>
    )
}

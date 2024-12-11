import classes from "./items-list.module.scss"
import { productsMaper } from "../lib/maper"
import { Product } from "@entities/product"
import { Spinner } from "@nextui-org/react"

type ItemsListProps = {
    data: Product[] | undefined
}

export function ItemsList({ data }: ItemsListProps) {
    // const { data } = useGetProductsQuery()

    return (
        <div className={classes.list}>
            {data ? productsMaper(data) : <Spinner label='loading...' />}
        </div>
    )
}

import { item, Product } from "@entities/product"
import { AddToBusket } from "@features/add-to-basket"
import { Dispatch, SetStateAction } from "react"

type usableProductProps = {
    product: item
    globalConnector: Dispatch<SetStateAction<item[]>>
}

export function UsableProduct({
    product,
    globalConnector,
}: usableProductProps) {
    return (
        <Product name={product.name} price={product.price}>
            <AddToBusket product={product} globalConnector={globalConnector} />
        </Product>
    )
}

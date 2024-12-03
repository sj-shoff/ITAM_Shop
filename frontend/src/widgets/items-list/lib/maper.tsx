import { Product } from "@entities/product"
import { UsableProduct } from "@widgets/usable-product"
import { ReactNode } from "react"

export function productsMaper(products: Product[]): ReactNode {
    const data = products.map((el) => (
        <UsableProduct key={el.id} product={el} />
    ))

    return data
}

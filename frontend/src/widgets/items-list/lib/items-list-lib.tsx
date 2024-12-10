import { item } from "@entities/product"
import { UsableProduct } from "@widgets/usable-product"
import { Dispatch, ReactElement, SetStateAction } from "react"

export function itemsMapper(
    data: item[],
    globalConnector: Dispatch<SetStateAction<item[]>>
): ReactElement[] {
    const res = data.map((el) => (
        <UsableProduct
            key={el.id}
            product={el}
            globalConnector={globalConnector}
        />
    ))

    return res
}

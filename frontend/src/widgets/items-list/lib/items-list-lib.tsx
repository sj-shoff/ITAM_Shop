import { getItems, Product } from "@entities/product"
import { ReactElement } from "react"

export async function itemsMapper(): Promise<ReactElement[]> {
    const data = await getItems()

    const res = data.map((el) => (
        <Product
            key={el.id}
            name={el.name}
            price={el.price}
            onClick={() => {}}
        />
    ))

    return res
}

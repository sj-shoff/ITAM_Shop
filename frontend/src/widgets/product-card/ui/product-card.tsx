import { useGetProductQuery } from "@entities/product"
import { useParams } from "react-router-dom"

export function ProductCard() {
    const id = useParams()
    console.log(id)
    const { data } = useGetProductQuery(id.product_id)

    return <></>
}

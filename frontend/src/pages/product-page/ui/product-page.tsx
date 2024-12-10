import { ProductId } from "@entities/product"
import { ProductCard } from "@widgets/product-card"
import { useParams } from "react-router-dom"

export function ProductPage() {
    const id = useParams()
    console.log(id)

    return (
        <section>
            <ProductCard />
        </section>
    )
}

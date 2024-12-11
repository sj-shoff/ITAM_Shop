import { useGetWishlistQuery } from "@entities/product"
import { ItemsList } from "@widgets/items-list"

export function WishlistPage() {
    const { data } = useGetWishlistQuery()

    return (
        <section>
            <ItemsList data={data} />
        </section>
    )
}

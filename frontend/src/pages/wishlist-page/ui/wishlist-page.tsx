import { useGetWishlistQuery } from "@entities/product"
import { ItemsList } from "@widgets/items-list"
import classes from "./wishlist-page.module.scss"

export function WishlistPage() {
    const { data } = useGetWishlistQuery()

    return (
        <section className={classes.wishlist}>
            <h1 className='text-7xl'>Список желаемого</h1>
            <ItemsList data={data} />
        </section>
    )
}

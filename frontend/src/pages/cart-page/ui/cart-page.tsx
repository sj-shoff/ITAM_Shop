import { useGetCartQuery } from "@entities/product"
import { ItemsList } from "@widgets/items-list"
import classes from "./cart-page.module.scss"

export function CartPage() {
    const { data } = useGetCartQuery()

    return (
        <section className={classes.cart}>
            <h1 className='text-7xl'>Корзина</h1>
            <ItemsList data={data} />
        </section>
    )
}

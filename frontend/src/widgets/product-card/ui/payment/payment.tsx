import { ProductId, ProductPrice } from "@entities/product"
import { ProductStockQuantity } from "@entities/product"
import { Button } from "@nextui-org/react"
import classes from "./payment.module.scss"
import { AddToCart } from "@features/add-to-cart"
import { AddToWishList } from "@features/add-to-wishlist"

type PaymentProps = {
    productId: ProductId | undefined
    price: ProductPrice
    quantity: ProductStockQuantity
}

export function Payment({ productId = 0, price, quantity }: PaymentProps) {
    console.log(productId)

    return (
        <div className={classes.payment}>
            <h2 className={classes.price}>Цена: {price} ₽</h2>
            <div className={classes.paymentContent}>
                <div className={classes.quantityBlock}>
                    <span className={classes.label}>В наличии: </span>
                    <span className={classes.quantity}>{quantity}</span>
                </div>
                <div className={classes.buttonsLayout}>
                    <Button fullWidth color='primary'>
                        Купить
                    </Button>
                    <AddToCart isIconOnly productId={productId} />
                    <AddToWishList isIconOnly productId={productId} />
                </div>
            </div>
        </div>
    )
}

import { ProductPrice } from "@entities/product"
import { ProductStockQuantity } from "@entities/product"
import { Button } from "@nextui-org/react"
import classes from "./payment.module.scss"
import { BasketIcon, PacketIcon } from "@shared/ui/icons"

type PaymentProps = ProductStockQuantity

export function Payment({
    price,
    quantity,
}: {
    price: ProductPrice
    quantity: PaymentProps
}) {
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
                    <Button isIconOnly>
                        <BasketIcon />
                    </Button>
                    <Button isIconOnly>
                        <PacketIcon />
                    </Button>
                </div>
            </div>
        </div>
    )
}

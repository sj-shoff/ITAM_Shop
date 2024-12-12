import { ProductId } from "@entities/product"
import { useRemoveFromCartMutation } from "@entities/product"
import { Button } from "@shared/ui/button"
import { MouseEventHandler } from "react"
import classes from "./remove-from-cart.module.scss"
import { RemoveIcon } from "@shared/ui/icons"

type addToBusketProps = {
    productId: ProductId
    isSmall?: boolean
}

export function RemoveFromCart({
    productId,
    isSmall = false,
}: addToBusketProps) {
    const [removeFromCart] = useRemoveFromCartMutation()

    const clickHandler: MouseEventHandler<HTMLButtonElement> = (e) => {
        e.stopPropagation()
        removeFromCart(productId)
        console.log(`id:${productId} added to the buset`)
    }

    return (
        <Button className={classes.removeFromCart} onClick={clickHandler}>
            <RemoveIcon className={isSmall && classes.small} />
            <span>Удалить из корзины</span>
        </Button>
    )
}

import { ProductId, useRemoveFromWishListMutation } from "@entities/product"
import { Button } from "@shared/ui/button"
import { RemoveIcon } from "@shared/ui/icons"
import { MouseEventHandler } from "react"
import classes from "./remove-from-wishlist.module.scss"

type addToBusketProps = {
    productId: ProductId
}

export function RemoveFromWishlist({ productId }: addToBusketProps) {
    const [removeFromWishlist] = useRemoveFromWishListMutation()

    const clickHandler: MouseEventHandler<HTMLButtonElement> = (e) => {
        e.stopPropagation()
        removeFromWishlist(productId)
        console.log(`id:${productId} added to the buset`)
    }

    return (
        <Button
            className={classes.removeFromWishlist}
            isIconOnly
            onClick={clickHandler}
        >
            <RemoveIcon />
        </Button>
    )
}

import { ProductId, useRemoveFromWishListMutation } from "@entities/product"
import { Button } from "@shared/ui/button"
import { RemoveIcon } from "@shared/ui/icons"
import { MouseEventHandler } from "react"

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
        <Button isIconOnly onClick={clickHandler}>
            <RemoveIcon />
        </Button>
    )
}

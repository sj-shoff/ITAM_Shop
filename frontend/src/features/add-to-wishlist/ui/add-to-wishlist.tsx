import { Button } from "@shared/ui/button"
import { WishListIcon } from "@shared/ui/icons"
import classes from "./add-to-wishlist.module.scss"
import { ProductId } from "@entities/product"
import { MouseEventHandler, useState } from "react"
import { useAddToWishlistMutation } from "@entities/product"

type AddToWishlistProps = {
    productId: ProductId
}

export function AddToWishList({ productId }: AddToWishlistProps) {
    // const [] = useAddToWishListMutation()
    const [isClicked, setIsClicked] = useState<boolean>(false)
    const [addToWishlist] = useAddToWishlistMutation()

    const clickHandler: MouseEventHandler<HTMLButtonElement> = (e) => {
        e.stopPropagation()
        setIsClicked(!isClicked)
        addToWishlist(productId)
    }

    return (
        <Button
            className={`${classes.addToWishlist} ${isClicked ? classes.clicked : classes.basic}`}
            isIconOnly
            onClick={clickHandler}
        >
            <WishListIcon />
        </Button>
    )
}

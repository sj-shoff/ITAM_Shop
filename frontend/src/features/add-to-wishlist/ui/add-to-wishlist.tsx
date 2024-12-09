import { Button } from "@shared/ui/button"
import { WishListIcon } from "@shared/ui/icons"
import classes from "./add-to-wishlist.module.scss"
import { ProductId } from "@entities/product"
import { useState } from "react"

type AddToWishlistProps = {
    productId: ProductId
}

export function AddToWishList({ productId }: AddToWishlistProps) {
    const [isClicked, setIsClicked] = useState<boolean>(false)

    function clickHandler() {
        console.log("123")
        setIsClicked(!isClicked)
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

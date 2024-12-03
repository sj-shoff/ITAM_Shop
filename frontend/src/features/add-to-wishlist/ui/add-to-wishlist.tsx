import { Button } from "@shared/ui/button"
import { WishListIcon } from "@shared/ui/icons"
import { productId } from "@entities/product"
import classes from "./add-to-wishlist.module.scss"

type AddToWishlistProps = {
    productId: productId
}

export function AddToWishList({ productId }: AddToWishlistProps) {
    function clickHandler() {
        console.log(productId)
    }

    return (
        <Button
            className={classes.addToWishlist}
            isIconOnly
            onClick={clickHandler}
        >
            <WishListIcon />
        </Button>
    )
}

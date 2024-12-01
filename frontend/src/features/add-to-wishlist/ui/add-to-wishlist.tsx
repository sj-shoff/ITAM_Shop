import { Button } from "@shared/ui/button"
import { WishListIcon } from "@shared/ui/icons"
import { useState } from "react"
import classes from "./add-to-wishlist.module.scss"

export function AddToWishList() {
    const [classValue, setClassValue] = useState(classes.basic)

    function clickHandler() {
        console.log("click")
    }

    return (
        <Button
            className={`${classes.addToWishlist} ${classValue}`}
            isIconOnly
            onClick={clickHandler}
        >
            <WishListIcon />
        </Button>
    )
}

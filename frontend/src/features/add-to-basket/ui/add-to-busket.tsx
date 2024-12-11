import { ProductId, useAddToCartMutation } from "@entities/product"
import { Button } from "@shared/ui/button"
import { MouseEventHandler } from "react"

type addToBusketProps = {
    productId: ProductId
}

export function AddToBusket({ productId }: addToBusketProps) {
    const [addToCart] = useAddToCartMutation()

    const clickHandler: MouseEventHandler<HTMLButtonElement> = (e) => {
        e.stopPropagation()
        addToCart(productId)
        console.log(`id:${productId} added to the buset`)
    }

    return <Button onClick={clickHandler}>Добавить в корзину</Button>
}

import { ProductId } from "@entities/product"
import { useRemoveFromCartMutation } from "@entities/product"
import { Button } from "@shared/ui/button"
import { MouseEventHandler } from "react"

type addToBusketProps = {
    productId: ProductId
}

export function RemoveFromCart({ productId }: addToBusketProps) {
    const [removeFromCart] = useRemoveFromCartMutation()

    const clickHandler: MouseEventHandler<HTMLButtonElement> = (e) => {
        e.stopPropagation()
        removeFromCart(productId)
        console.log(`id:${productId} added to the buset`)
    }

    return <Button onClick={clickHandler}>Удалить из корзины</Button>
}

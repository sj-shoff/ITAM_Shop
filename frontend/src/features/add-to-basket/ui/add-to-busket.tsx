import { ProductId } from "@entities/product"
import { Button } from "@shared/ui/button"
import { MouseEventHandler } from "react"

type addToBusketProps = {
    productId: ProductId
}

export function AddToBusket({ productId }: addToBusketProps) {
    const clickHandler: MouseEventHandler<HTMLButtonElement> = (e) => {
        e.stopPropagation()
        console.log(productId)
    }

    return <Button onClick={clickHandler}>Добавить в корзину</Button>
}

import { Product, productId } from "@entities/product"
import { Button } from "@shared/ui/button"

type addToBusketProps = {
    productId: productId
}

export function AddToBusket({ productId }: addToBusketProps) {
    function clickHandler() {
        // TODO: add logic
    }

    return <Button onClick={clickHandler}>Добавить в корзину</Button>
}

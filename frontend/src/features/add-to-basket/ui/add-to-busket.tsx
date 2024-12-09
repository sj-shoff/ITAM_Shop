import { ProductId } from "@entities/product"
import { Button } from "@shared/ui/button"

type addToBusketProps = {
    productId: ProductId
}

export function AddToBusket({ productId }: addToBusketProps) {
    function clickHandler() {
        console.log(productId)
    }

    return <Button onClick={clickHandler}>Добавить в корзину</Button>
}

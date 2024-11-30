import { item } from "@entities/product"
import { Button } from "@nextui-org/react"
import { Dispatch, SetStateAction } from "react"

type addToBusketProps = {
    product: item
    globalConnector: Dispatch<SetStateAction<item[]>>
}

export function AddToBusket({ product, globalConnector }: addToBusketProps) {
    function clickHandler() {
        globalConnector((prev) => prev.filter((el) => el.id === product.id))
    }

    return <Button onClick={clickHandler}>Добавить в корзину</Button>
}

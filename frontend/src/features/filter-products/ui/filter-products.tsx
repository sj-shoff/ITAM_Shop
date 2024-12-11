import { FilterOptions } from "@entities/product"
import { Button } from "@nextui-org/react"
import { MouseEventHandler } from "react"

type FilterProductsProps = {
    clickHandler: MouseEventHandler<HTMLButtonElement>
    filterOptions: FilterOptions
}

export function FilterProducts({ clickHandler, ...rest }: FilterProductsProps) {
    return (
        <Button {...rest} onClick={clickHandler}>
            отфильтровать
        </Button>
    )
}

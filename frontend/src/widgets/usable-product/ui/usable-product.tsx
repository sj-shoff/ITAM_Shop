import { ProductComponent, Product, productId } from "@entities/product"
import { AddToBusket } from "@features/add-to-basket"
import { AddToWishList } from "@features/add-to-wishlist"
import { useAppSelector } from "@shared/lib/state"


type usableProductProps = {
    id: productId
}

export function UsableProduct({ id }: usableProductProps) {
    const product = useAppSelector((state) => state.catalog.products)

    return (
        <ProductComponent
            headContent={<AddToWishList productId={id} />}
            {.}
        >
            <AddToBusket productId={id} />
        </ProductComponent>
    )
}

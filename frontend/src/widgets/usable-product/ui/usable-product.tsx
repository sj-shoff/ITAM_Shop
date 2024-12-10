import { Product, ProductComponent, ProductId } from "@entities/product"
import { AddToBusket } from "@features/add-to-basket"
import { AddToWishList } from "@features/add-to-wishlist"

type usableProductProps = {
    id: ProductId
    product: Partial<Product>
}

export function UsableProduct({ id, product }: usableProductProps) {
    return (
        <ProductComponent
            product={product}
            headContent={<AddToWishList productId={id} />}
        >
            <AddToBusket productId={id} />
        </ProductComponent>
    )
}

import { Product, ProductComponent, ProductId } from "@entities/product"
import { AddToCart } from "@features/add-to-cart"
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
            <AddToCart productId={id} />
        </ProductComponent>
    )
}

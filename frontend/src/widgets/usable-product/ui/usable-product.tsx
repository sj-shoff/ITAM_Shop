import { Product, ProductComponent, ProductId } from "@entities/product"
import { AddToCart } from "@features/add-to-cart"
import { AddToWishList } from "@features/add-to-wishlist"
import { RemoveFromCart } from "@features/remove-from-cart"
import { RemoveFromWishlist } from "@features/remove-from-wishlist copy"

type usableProductProps = {
    id: ProductId
    product: Partial<Product>
}

export function UsableProduct({ id, product }: usableProductProps) {
    return (
        <ProductComponent
            product={product}
            headContent={
                product.is_in_fav === 1 ? (
                    <RemoveFromWishlist productId={id} />
                ) : (
                    <AddToWishList productId={id} />
                )
            }
        >
            {product.is_in_cart === 1 ? (
                <RemoveFromCart productId={id} isSmall />
            ) : (
                <AddToCart productId={id} />
            )}
        </ProductComponent>
    )
}

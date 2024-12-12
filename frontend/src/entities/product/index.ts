export { ProductComponent } from "./ui/product"
export {
    productsApi,
    useGetProductsQuery,
    useGetProductQuery,
    useAddToCartMutation,
    useGetCartQuery,
    useGetWishlistQuery,
    useRemoveFromCartMutation,
    useAddToWishlistMutation,
    useRemoveFromWishListMutation,
} from "./api/item-api"
export type {
    Product,
    ProductId,
    ProductName,
    ProductCategory,
    ProductDescription,
    ProductPrice,
    ProductStockQuantity,
    ProductImage,
    ProductIsInCart,
    ProductIsInFav,
    FilterOptions,
    FilterCategory,
    FilterMaxPrice,
    FilterMinPrice,
    Feature,
    FeatureName,
    FeatureType,
} from "./model/product-model"
export {
    featureDefaultValue,
    productDefaultValue,
    requestTypeDefaultVale,
} from "./model/product-model"
export { productSlice, getList } from "./model/product-slice"

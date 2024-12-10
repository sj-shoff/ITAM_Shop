export { ProductComponent } from "./ui/product"
export {
    productsApi,
    useGetProductsQuery,
    useGetProductQuery,
    useAddToWishListMutation,
    useFilterProductsMutation,
} from "./api/item-api"
export type {
    Product,
    ProductId,
    ProductName,
    ProductCategory,
    ProductDescription,
    ProductImageUrl,
    ProductPrice,
    ProductSpecifications,
    ProductStockQuantity,
    FilterOptions,
    FilterCategory,
    FilterMaxPrice,
    FilterMinPrice,
} from "./model/product-model"
export { productSlice } from "./model/product-slice"

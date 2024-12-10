export { ProductComponent } from "./ui/product"
export {
    productsApi,
    useGetProductsQuery,
    useGetProductQuery,
    useAddToWishListMutation,
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
} from "./model/product-model"
export { productSlice } from "./model/product-slice"

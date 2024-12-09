export { ProductComponent } from "./ui/product"
export {
    productsApi,
    useGetProductsQuery,
    useGetProductQuery,
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
} from "./model/product-model"
export { productSlice } from "./model/product-slice"

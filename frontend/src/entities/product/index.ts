export { ProductComponent } from "./ui/product"
export {
    productsApi,
    useGetProductsQuery,
    useGetProductQuery,
    useFilterProductsMutation,
} from "./api/item-api"
export type {
    Product,
    ProductId,
    ProductName,
    ProductCategory,
    ProductDescription,
    ProductPrice,
    ProductStockQuantity,
    FilterOptions,
    FilterCategory,
    FilterMaxPrice,
    FilterMinPrice,
} from "./model/product-model"
export { productSlice } from "./model/product-slice"

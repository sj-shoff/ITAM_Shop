export { ProductComponent } from "./ui/product"
export { productsApi, useGetProductsQuery } from "./api/item-api"
export {
    type Product,
    type ProductId,
    type ProductName,
    type ProductCategory,
    type ProductDescription,
    type ProductImageUrl,
    type ProductPrice,
    type ProductSpecifications,
} from "./model/product-model"
export { productSlice } from "./model/product-slice"

import { baseApi } from "@shared/api"
import {
    requestDTOschema,
    Product,
    productDTOschema,
    ProductId,
    RequestType,
} from "../model/product-model"

export const productsApi = baseApi.injectEndpoints({
    endpoints: (create) => ({
        // GET Product
        getProduct: create.query<RequestType, ProductId>({
            query: (productId) => `/get_item_page/${productId}`,
            transformResponse: (responce: unknown) =>
                requestDTOschema.parse(responce),
            providesTags: ["Product"],
        }),

        // GET Сatalog
        getProducts: create.query<Product[], void>({
            query: () => "/catalog",
            transformResponse: (responce: unknown) =>
                productDTOschema.array().parse(responce),
            providesTags: ["Catalog"],
        }),
    }),
    overrideExisting: true,
})

export const {
    // Product
    useGetProductQuery,

    // Catalog
    useGetProductsQuery,
} = productsApi

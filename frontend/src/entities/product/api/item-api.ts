import { baseApi } from "@shared/api"
import { z } from "zod"
import { Product, ProductId } from "../model/product-model"

const productDTOschema = z.object({
    product_id: z.number(),
    product_price: z.number(),
    product_name: z.string(),
    product_image: z.string(),
    product_description: z.string(),
    product_specifications: z.string(),
    product_category: z.string(),
    product_stock_quantity: z.number(),
})

export const productsApi = baseApi.injectEndpoints({
    endpoints: (create) => ({
        // Сatalog
        getProducts: create.query<Product[], void>({
            query: () => "/catalog",
            transformResponse: (responce: unknown) =>
                productDTOschema.array().parse(responce),
            providesTags: ["Catalog"],
        }),
        getProduct: create.query<Product, ProductId>({
            query: (productId) => `/get_item_page/${productId}`,
            transformResponse: (responce: unknown) =>
                productDTOschema.parse(responce),
            providesTags: ["Product"],
        }),

        // Wishlist
        getWishlist: create.query<Product[], void>({
            query: () => "/fav_items",
            transformResponse: (responce: unknown) =>
                productDTOschema.array().parse(responce),
            providesTags: ["Wishlist"],
        }),
        getWishlistItem: create.query<Product, ProductId>({
            query: (productId) => `/fav_items/${productId}`,
            transformResponse: (responce: unknown) =>
                productDTOschema.parse(responce),
            providesTags: ["Product"],
        }),
        addToWishlist: create.mutation<void, ProductId>({
            query: (id) => ({ method: "Post", url: `fav_items/${id}` }),
        }),
    }),
    overrideExisting: true,
})

export const { useGetProductsQuery, useGetProductQuery, useGetWishlistQuery } =
    productsApi

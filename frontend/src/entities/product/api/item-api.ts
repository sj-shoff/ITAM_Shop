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
        addToWishlist: create.mutation<void, ProductId>({
            query: (id) => ({ method: "POST", url: `fav_items/${id}` }),
            // Обновляем все списки товаров и отображаем новый ui
            invalidatesTags: ["Cart", "Catalog", "Wishlist"],
        }),
        removeFromWishList: create.mutation<void, ProductId>({
            query: (id) => ({ method: "DELETE", url: `/fav_items/${id}` }),
            // Перерисовываем все списки и страницу с товаром,
            // чтобы отобразить ui
            invalidatesTags: (result, error, productId) => [
                "Cart",
                "Catalog",
                "Wishlist",
                { type: "Product", id: productId },
            ],
        }),

        // Cart
        getCart: create.query<Product[], void>({
            query: () => "/cart",
            transformResponse: (responce: unknown) =>
                productDTOschema.array().parse(responce),
            providesTags: ["Cart"],
        }),
        addToCart: create.mutation<void, ProductId>({
            query: (id) => ({ method: "POST", url: `/cart/add/${id}` }),
            // Обновляем во всех списках, что теперь товар в корзине
            // В том числе перерисовываем страницу с определенным товаром, чтобы отобразить новый ui
            invalidatesTags: (result, error, productId) => [
                "Cart",
                "Catalog",
                "Wishlist",
                { type: "Product", id: productId },
            ],
        }),
        removeFromCart: create.mutation<void, ProductId>({
            query: (id) => ({ method: "DELETE", url: `/cart/remove/${id}` }),
            // Обновляем во всех списках, что теперь товар НЕ в корзине
            invalidatesTags: (result, error, productId) => [
                "Cart",
                "Catalog",
                "Wishlist",
                { type: "Product", id: productId },
            ],
        }),
    }),
    overrideExisting: true,
})

export const {
    useGetProductQuery,
    // Catalog
    useGetProductsQuery,

    // Wishlist
    useGetWishlistQuery,
    useAddToWishlistMutation,
    useRemoveFromWishListMutation,

    // Cart
    useGetCartQuery,
    useAddToCartMutation,
    useRemoveFromCartMutation,
} = productsApi

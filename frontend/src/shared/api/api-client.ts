import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react"

const baseUrl = "http://147.45.163.58:8080/"

export const baseApi = createApi({
    baseQuery: fetchBaseQuery({ baseUrl }),
    endpoints: () => ({}),
    tagTypes: [
        "Catalog",
        "Product",
        "Wishlist",
        // Продукты внутри вишлиста
        "WishlistItem",
        "Cart",
        // Продукты внутри корзины
        "CartItem",
    ],
})

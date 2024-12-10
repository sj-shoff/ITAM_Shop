import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react"

const baseUrl = "http://localhost:3000"

export const baseApi = createApi({
    baseQuery: fetchBaseQuery({ baseUrl }),
    endpoints: () => ({}),
    tagTypes: ["Catalog", "Product", "Wishlist"],
})

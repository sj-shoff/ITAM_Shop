import { baseApi } from "@shared/api"
import { Product, productDTOschema, ProductId } from "../model/product-model"

export const wishListApi = baseApi.injectEndpoints({
    endpoints: (create) => ({
        // Wishlist
        // GET
        getWishlist: create.query<Product[], void>({
            query: () => "/fav_items",
            transformResponse: (responce: unknown) =>
                productDTOschema.array().parse(responce),
            providesTags: ["Wishlist"],
        }),

        // POST
        addToWishlist: create.mutation<void, ProductId>({
            query: (id) => ({ method: "POST", url: `fav_items/${id}` }),
            // Обновляем все списки товаров
            // и отображаем новый ui на странице товара
            invalidatesTags: (result, error, productId) => [
                "Cart",
                "Catalog",
                "Wishlist",
                { type: "Product", id: productId },
            ],
        }),

        // DELETE
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
    }),
    overrideExisting: true,
})

export const {
    useAddToWishlistMutation,
    useGetWishlistQuery,
    useRemoveFromWishListMutation,
} = wishListApi

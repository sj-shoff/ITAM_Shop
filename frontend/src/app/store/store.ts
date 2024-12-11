import { productsApi } from "@entities/product"
import { configureStore } from "@reduxjs/toolkit"
import { productSlice } from "@entities/product"

export const store = configureStore({
    reducer: {
        [productsApi.reducerPath]: productsApi.reducer,
        [productSlice.name]: productSlice.reducer,
    },
    middleware: (getDefaultMiddleware) =>
        getDefaultMiddleware().concat(productsApi.middleware),
})

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch

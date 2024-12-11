import { createSlice, PayloadAction } from "@reduxjs/toolkit"
import { FilterOptions, Product } from "./product-model"

type ProductState = {
    filteredList: Product[]
}

const initStateValue: ProductState = {
    filteredList: [
        {
            product_id: 0,
            product_category: "",
            product_image: "",
            product_description: "",
            product_name: "",
            product_price: 0,
            product_stock_quantity: 0,
        },
    ],
}

export const productSlice = createSlice({
    name: "products",
    initialState: initStateValue,
    reducers: {
        // Обновление стейта
        insertProducts(state, action: PayloadAction<Product[]>) {
            state.filteredList = action.payload
        },
        // Фильтрация
        filterProducts(state, action: PayloadAction<FilterOptions>) {
            state.filteredList = []
        },
    },
    selectors: {
        getList: (state) => state.filteredList,
    },
})

export const { getList } = productSlice.selectors

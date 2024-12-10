import { createSlice, PayloadAction } from "@reduxjs/toolkit"
import { Product, productId, productName, productPrice } from "./product-model"

type ProductState = Product

const initStateValue: ProductState = {
    id: 1,
    name: "Кожанка",
    price: 1000,
}

export const productSlice = createSlice({
    name: "product",
    initialState: initStateValue,
    reducers: {
        changeId(state, action: PayloadAction<productId>) {
            state.id = action.payload
        },
        changePrice(state, action: PayloadAction<productPrice>) {
            state.price = action.payload
        },
        changeName(state, action: PayloadAction<productName>) {
            state.name = action.payload
        },
    },
    selectors: {
        getPrice: (state) => state.price,
        getName: (state) => state.name,
        getId: (state) => state.id,
    },
})

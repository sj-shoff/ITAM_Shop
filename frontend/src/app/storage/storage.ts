import { Product, productId } from "@entities/product"
import { configureStore, Reducer } from "@reduxjs/toolkit"

type state = {
    jwt: string
    registation: boolean
    products: Product[]
    catalog: {
        products: Record<productId, Product>
        ids: number[]
    }
}

type FetchDataAction = {
    type: "fetch products"
    payload: Product[]
}

type RemoveFromCatalogAction = {
    type: "remove item from catalog"
    payload: {
        id: productId
    }
}

type action = FetchDataAction | RemoveFromCatalogAction

const initialstate: state = {
    jwt: "",
    registation: false,
    products: [],
    catalog: {
        products: [],
        ids: [],
    },
}
// Загружает данные из инициализации
const initialProducts: Product[] = [
    {
        name: "Худи",
        price: 1000,
        id: 1,
    },
    {
        name: "Штаны",
        price: 1000,
        id: 2,
    },
    {
        name: "Шапка",
        price: 1000,
        id: 3,
    },
]
const initialCatalog: Product[] = initialProducts

const reducer: Reducer<state, action, state> = (
    state = initialstate,
    action
) => {
    switch (action.type) {
        case "fetch products": {
            return {
                ...state,
                products: action.payload ?? initialProducts,
            }
        }
        case "remove item from catalog": {
            const { id } = action.payload
            const prevCatalog = { ...state.catalog }

            if (prevCatalog) {
                delete prevCatalog.ids[id]
                delete prevCatalog.products[id]
            }

            return {
                ...state,
                catalog: {
                    ...state.catalog,
                    products: {},
                },
            }
        }
        default: {
            console.log("encountered unknown action")
            return state
        }
    }
}

export const store = configureStore({
    reducer: reducer,
})

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch

import { router } from "./router/router"
import { Provider } from "react-redux"
import { store } from "./store"
import "./main.scss"
import { RouterProvider } from "react-router-dom"

export function App() {
    return (
        <Provider store={store}>
            <RouterProvider router={router} />
        </Provider>
    )
}

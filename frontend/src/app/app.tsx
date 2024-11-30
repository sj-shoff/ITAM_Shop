import { RouterProvider } from "react-router-dom"

import "./main.scss"
import { router } from "./router/router"
import { Provider } from "react-redux"
import { store } from "./storage"

export function App() {
    return (
        <Provider store={store}>
            <RouterProvider router={router} />
        </Provider>
    )
}

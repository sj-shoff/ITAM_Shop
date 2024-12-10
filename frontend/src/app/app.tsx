import { router } from "./router/router"
import { Provider } from "react-redux"
import { store } from "./store"
import "./main.scss"
import { RouterProvider } from "react-router-dom"

export function App() {
    return (
        <div className='dark text-foreground w-full h-full flex justify-center'>
            <Provider store={store}>
                <RouterProvider router={router} />
            </Provider>
        </div>
    )
}

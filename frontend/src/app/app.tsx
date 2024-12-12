import { router } from "./router/router"
import { Provider } from "react-redux"
import { store } from "./store"
import "./main.scss"
import { RouterProvider } from "react-router-dom"
import { NextUIProvider } from "@nextui-org/react"

export function App() {
    return (
        <div className='dark text-foreground w-full h-full flex justify-center'>
            <Provider store={store}>
                <NextUIProvider className='w-full h-full dark flex flex-col items-center'>
                    <RouterProvider router={router} />
                </NextUIProvider>
            </Provider>
        </div>
    )
}

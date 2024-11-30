import { Homepage } from "@pages/home-page"
import { LoginPage } from "@pages/login"
import { createBrowserRouter } from "react-router-dom"

export const router = createBrowserRouter([
    {
        path: "/",
        element: <Homepage />,
        children: [
            {
                element: <LoginPage />,
                path: "/login",
            },
        ],
    },
])

import { CatalogPage, Homepage } from "@pages/catalog-page"
import { LoginPage } from "@pages/login"
import { ReactElement } from "react"
import { createBrowserRouter, Route, Routes } from "react-router-dom"

type link = {
    link: string
    element: ReactElement
}

const routesConfig: link[] = [
    {
        link: "/",
        element: <Homepage />,
    },
    {
        link: "/login",
        element: <LoginPage />,
    },
]

export function AppRoutes() {
    return (
        <Routes>
            {routesConfig.map((el) => (
                <Route key={el.link} path={el.link} element={el.element} />
            ))}
        </Routes>
    )
}

export const router = createBrowserRouter([
    {
        path: "/",
        element: <Homepage />,
        children: [
            {
                path: "/catalog",
                element: <CatalogPage />,
            },
        ],
    },
])

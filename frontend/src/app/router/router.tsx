import { Spinner } from "@nextui-org/react"
import { DefaultLayout } from "../layouts/dafault-layout/ui/dafault-layout"
import { CatalogPage } from "@pages/catalog-page"
import { Homepage } from "@pages/home-page"
import { createBrowserRouter } from "react-router-dom"
import { ProductPage } from "@pages/product-page"
import { WishlistPage } from "@pages/wishlist-page"
import { CartPage } from "@pages/cart-page"

export const router = createBrowserRouter([
    {
        path: "/",
        element: <DefaultLayout />,
        errorElement: (
            <Spinner
                style={{ marginTop: "7rem" }}
                label='Oups... Application crashed or this page doesn`t exist'
                color='danger'
                labelColor='danger'
            />
        ),
        children: [
            {
                path: "/",
                element: <Homepage />,
                errorElement: (
                    <Spinner
                        style={{ marginTop: "7rem" }}
                        label='Oups... Can`t access to the homepage'
                        color='danger'
                        labelColor='danger'
                    />
                ),
            },
            {
                path: "catalog",
                element: <CatalogPage />,
                errorElement: (
                    <Spinner
                        style={{ marginTop: "7rem" }}
                        label='somthing went wrong on catalog page...'
                        color='danger'
                        labelColor='danger'
                    />
                ),
            },
            {
                path: "catalog/:product_id",
                element: <ProductPage />,
            },
            {
                path: "fav_items",
                element: <WishlistPage />,
            },
            {
                path: "fav_items/:product_id",
                element: <ProductPage />,
            },
            {
                path: "cart",
                element: <CartPage />,
            },
        ],
    },
])

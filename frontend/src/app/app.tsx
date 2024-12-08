import { AppRoutes } from "./router/router"
import { Provider } from "react-redux"
import { store } from "./store"
import { DefaultLayout } from "./layouts/dafault-layout/dafault-layout"
import "./main.scss"

export function App() {
    return (
        <Provider store={store}>
            <DefaultLayout>
                <AppRoutes />
            </DefaultLayout>
        </Provider>
    )
}

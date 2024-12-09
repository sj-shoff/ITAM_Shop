import { AppRoutes } from "./router/router"
import { Provider } from "react-redux"
import { store } from "./store"
import { DefaultLayout } from "./layouts/dafault-layout/ui/dafault-layout"
import "./main.scss"

export function App() {
    return (
        <Provider store={store}>
            <div className='dark'>
                <DefaultLayout>
                    <AppRoutes />
                </DefaultLayout>
            </div>
        </Provider>
    )
}

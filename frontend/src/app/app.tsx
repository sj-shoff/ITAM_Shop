import { BrowserRouter, Routes, Route } from "react-router-dom"
import { Homepage } from "@pages/home-page"
import "./index.scss"

export function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path={"/"} element={<Homepage />} />
            </Routes>
        </BrowserRouter>
    )
}

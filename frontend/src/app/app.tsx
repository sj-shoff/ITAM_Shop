import { Routes, Route } from "react-router-dom"
import { Homepage } from "@pages/home-page"
import Layout from "@pages/admin/ui/components/layout"
import "./main.scss"

export function App() {
    return (
        <Routes>
            <Route path={"/"} element={<Homepage />} />
        </Routes>
    )
}

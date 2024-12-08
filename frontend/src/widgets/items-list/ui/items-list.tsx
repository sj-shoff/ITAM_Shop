import classes from "./items-list.module.scss"
import { productsMaper } from "../lib/maper"
import { useSelector } from "react-redux"

export function ItemsList() {
    const catalog = useSelector()

    return <div className={classes.list}>{productsMaper(catalog)}</div>
}

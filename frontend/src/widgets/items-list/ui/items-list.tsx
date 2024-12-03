import { useAppSelector } from "@shared/lib"
import classes from "./items-list.module.scss"
import { productsMaper } from "../lib/maper"

export function ItemsList() {
    const catalog = useAppSelector((el) => el.catalog.ids)


    return <div className={classes.list}>{productsMaper()}</div>
}

import { ReactElement, useEffect, useState } from "react"
import classes from "./items-list.module.scss"
import { itemsMapper } from "../lib/items-list-lib"

export function ItemsList() {
    const [listState, setListState] = useState<ReactElement[]>()

    useEffect(() => {
        itemsMapper().then((value) => {
            setListState(value)
        })
    })

    return <div className={classes.list}>{listState}</div>
}

import { ReactElement, useEffect, useState } from "react"
import classes from "./items-list.module.scss"
import { itemsMapper } from "../lib/items-list-lib"
import { getItems, item } from "@entities/product"

export function ItemsList() {
    const [data, setData] = useState<item[] | null>()
    const [listState, setListState] = useState<ReactElement[] | null>()

    useEffect(() => {
        getItems().then((el) => {
            setData(el)
        })
    })

    useEffect(() => {
        if (data) {
            setListState(itemsMapper(data, () => {}))
        } else {
            // TODO: добавить обработчик ошибок
            console.log("[ERROR] no data")
            setListState(null)
        }
    }, [data])

    return <div className={classes.list}>{listState}</div>
}

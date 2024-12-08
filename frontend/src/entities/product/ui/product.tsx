import { ReactNode, useState } from "react"
import { Skeleton } from "@nextui-org/react"
import classes from "./prooduct.module.scss"
import { useSelector } from "react-redux"
import { productSlice } from "../model/product-slice"

type itemProps = {
    name: string
    price: number
    children: ReactNode
    headContent?: ReactNode
}

// TODO: refactor to the model
export type ProductUiElement = "image" | "name" | "price"

export function ProductComponent({
    name,
    price,
    children,
    headContent,
}: itemProps) {
    const [isLoaded, setIsLoaded] = useState<Record<ProductUiElement, boolean>>(
        {
            image: false,
            name: false,
            price: false,
        }
    )
    const name = useSelector(productSlice.selectors.getName)
    const id = useSelector(productSlice.selectors.getId)
    const price = useSelector(productSlice.selectors.getPrice)

    function imageLoadHandler() {
        setIsLoaded((prev) => ({
            ...prev,
            image: true,
        }))
    }

    return (
        <article className={classes.card}>
            <div className={classes.headContent}>{headContent}</div>
            <div className={classes.body}>
                <Skeleton isLoaded={isLoaded.image}>
                    <img
                        className={classes.image}
                        src='public/product-image-1.png'
                        alt='product image'
                        width='410'
                        height='460'
                        loading='lazy'
                        onLoad={imageLoadHandler}
                    />
                </Skeleton>
                <div>
                    <p className={classes.productInfo}>
                        <span className={classes.name}>{name}</span>
                        <span className={classes.price}>
                            <b>({price}₽)</b>
                        </span>
                    </p>
                </div>
                {/* Здесь будут все фичи */}
                {children}
            </div>
        </article>
    )
}

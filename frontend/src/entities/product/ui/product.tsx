import { ReactNode, useState } from "react"
import { Skeleton } from "@nextui-org/react"
import classes from "./prooduct.module.scss"

type itemProps = {
    name: string
    price: number
    children: ReactNode
}

// TODO: refactor to the model
export type ProductUiElement = "image" | "text-info"

export function Product({ name, price, children }: itemProps) {
    const [isLoaded, setIsLoaded] = useState<Record<ProductUiElement, boolean>>(
        {
            image: false,
            "text-info": false,
        }
    )

    function imageLoadHandler() {
        setIsLoaded((prev) => ({
            ...prev,
            image: true,
        }))
    }

    return (
        <article className={classes.card}>
            <div className={classes.body}>
                <Skeleton isLoaded={isLoaded.image ? true : false}>
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

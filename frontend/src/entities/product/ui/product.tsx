import { ReactNode, useState } from "react"
import { Skeleton } from "@nextui-org/react"
import classes from "./prooduct.module.scss"
import { Product } from "../model/product-model"

type itemProps = {
    product: Partial<Product>
    children: ReactNode
    headContent?: ReactNode
}

export function ProductComponent({
    product,
    headContent,
    children,
}: itemProps) {
    const [isImageLoaded, setIsImageLoaded] = useState<boolean>(false)
    const { product_name, product_image, product_price } = product

    function imageLoadHandler() {
        setIsImageLoaded(true)
    }

    return (
        <article className={classes.card}>
            <div className={classes.headContent}>{headContent}</div>
            <div className={classes.body}>
                <Skeleton isLoaded={isImageLoaded}>
                    <img
                        className={classes.image}
                        src={product_image}
                        alt='product image'
                        width='410'
                        height='460'
                        loading='lazy'
                        onLoad={imageLoadHandler}
                    />
                </Skeleton>
                <div>
                    <p className={classes.productInfo}>
                        <span className={classes.name}>{product_name}</span>
                        <span className={classes.price}>
                            <b>({product_price}₽)</b>
                        </span>
                    </p>
                </div>
                {/* Здесь будут все фичи */}
                {children}
            </div>
        </article>
    )
}

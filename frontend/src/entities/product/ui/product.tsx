import { ReactNode, useState } from "react"
import { Skeleton } from "@nextui-org/react"
import classes from "./prooduct.module.scss"
import { Product } from "../model/product-model"
import { useNavigate } from "react-router-dom"

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
    const navigate = useNavigate()
    const { product_name, product_image, product_price } = product

    function imageLoadHandler() {
        setIsImageLoaded(true)
    }

    function cardClickHandler() {
        navigate(`/catalog/${product.product_id}`)
    }

    return (
        <article className={classes.card} onClick={cardClickHandler}>
            <div className={classes.headContent}>{headContent}</div>
            <div className={classes.body}>
                <Skeleton
                    className={classes.imageContainer}
                    isLoaded={isImageLoaded}
                >
                    <img
                        className={classes.image}
                        src={`data:image/jpeg;base64,${product_image}`}
                        alt='product image'
                        width='410'
                        height='460'
                        loading='lazy'
                        onLoad={imageLoadHandler}
                    />
                </Skeleton>
                <div className={classes.productInfo}>
                    <span className={classes.name}>{product_name}</span>
                    <span className={classes.price}>
                        <b>({product_price}₽)</b>
                    </span>
                </div>
                {/* Здесь будут все фичи */}
                <div className={classes.features}>{children}</div>
            </div>
        </article>
    )
}

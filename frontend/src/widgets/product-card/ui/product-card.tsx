import { useGetProductQuery } from "@entities/product"
import { Spinner } from "@nextui-org/react"
import { useParams } from "react-router-dom"
import { Image } from "./image/image"
import classes from "./product-card.module.scss"
import { Payment } from "./payment/payment"
import { TextInfo } from "./text-info/text-info"
import { requestTypeDefaultVale } from "@entities/product"
import { Specifications } from "./specifications/specifications"

export function ProductCard() {
    // Получение id товара из роутинга
    const params = useParams<{ product_id: string }>()
    // Получение товара из стора или из запроса
    const { data = requestTypeDefaultVale } = useGetProductQuery(
        Number(params.product_id)
    )
    const { product, features } = data

    // Если данных нет, или они грузятся - отображается спинер
    if (!data) {
        return (
            <Spinner
                style={{ marginTop: "5rem" }}
                label='Loading...'
                color='warning'
                labelColor='warning'
            />
        )
    }

    return (
        <div className={classes.productCard}>
            <Image path={product.product_image} />
            <div className={classes.info}>
                <TextInfo
                    name={product.product_name}
                    category={product.product_category}
                    description={product.product_description}
                />
                <Payment
                    productId={Number(params.product_id)}
                    price={product.product_price}
                    quantity={product.product_stock_quantity}
                />
                <ul className={classes.specifications}>
                    <Specifications features={features} />
                </ul>
            </div>
        </div>
    )
}

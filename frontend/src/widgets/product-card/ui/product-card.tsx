import { useGetProductQuery } from "@entities/product"
import { Spinner } from "@nextui-org/react"
import { useParams } from "react-router-dom"
import { Image } from "./image/image"
import classes from "./product-card.module.scss"
import { Payment } from "./payment/payment"
import { TextInfo } from "./text-info/text-info"

export function ProductCard() {
    // Получение id товара из роутинга
    const params = useParams<{ product_id: string }>()
    // Получение товара из стора или из запроса
    const { data } = useGetProductQuery(Number(params.product_id))

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
            <Image path={data.product_image} />
            <div className={classes.info}>
                <TextInfo
                    name={data.product_name}
                    category={data.product_category}
                    description={data.product_description}
                />
                <Payment
                    price={data.product_price}
                    quantity={data.product_stock_quantity}
                />
                <div className={classes.specifications}>
                    {/* {data.product_specifications} */}
                    спецификации
                </div>
            </div>
        </div>
    )
}

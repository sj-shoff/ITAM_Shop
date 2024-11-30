import { ReactNode } from "react"
import classes from "./item.module.scss"
import { Card, CardBody, Skeleton } from "@nextui-org/react"

type itemProps = {
    name: string
    price: number
    children: ReactNode
}

export function Product({ name, price, children }: itemProps) {
    return (
        <Card>
            <CardBody>
                <Skeleton>
                    <img
                        className='class-for-img'
                        src='/'
                        alt=''
                        width='410'
                        height='460'
                        loading='lazy'
                    />
                </Skeleton>

                <div className={classes.header}>
                    <p>{name}</p>
                    <p>{price}₽</p>
                </div>
                {/* Здесь будут все фичи */}
                {children}
            </CardBody>
        </Card>
    )
}

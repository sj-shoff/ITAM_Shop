import { MouseEventHandler } from "react"
import classes from "./item.module.scss"
import { Button, Card, CardBody, Skeleton } from "@nextui-org/react"

type itemProps = {
    name: string
    price: number,
    onClick: MouseEventHandler<HTMLButtonElement>
}

export function Product({ name, price, onClick }: itemProps) {
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
                <Button onClick={onClick}>В корзину</Button>
            </CardBody>
        </Card>
    )
}

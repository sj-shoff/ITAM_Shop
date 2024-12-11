import { Button } from "@nextui-org/react"
import { ProductCard } from "@widgets/product-card"
import { useNavigate } from "react-router-dom"
import classes from "./product-page.module.scss"
import { BackIcon } from "@shared/ui/icons"

export function ProductPage() {
    const navigate = useNavigate()

    return (
        <section className={classes.productSection}>
            <Button
                className='text-lg'
                size='lg'
                variant='light'
                startContent={<BackIcon fill='white' />}
                onClick={() => {
                    // Возращает к прошлому роутингу
                    navigate(-1)
                }}
            >
                Назад
            </Button>
            {/* В дальнейшем можно расширить функционал страницы
            различными виджетами по типу - отзывы, рекомендации */}
            <ProductCard />
        </section>
    )
}

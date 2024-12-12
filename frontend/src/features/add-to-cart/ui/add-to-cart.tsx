import { ProductId, useAddToCartMutation } from "@entities/product"
import { Button } from "@shared/ui/button"
import {
    Button as ButtonNextUI,
    Divider,
    Modal,
    ModalBody,
    ModalContent,
    ModalHeader,
} from "@nextui-org/react"
import { PacketIcon } from "@shared/ui/icons"
import { MouseEventHandler, useEffect, useState } from "react"

type addToBusketProps = {
    productId: ProductId
    isIconOnly?: boolean
}

export function AddToCart({ productId, isIconOnly = false }: addToBusketProps) {
    const [addToCart, data] = useAddToCartMutation()
    const [isOpen, setIsOpen] = useState<boolean>(false)

    useEffect(() => {
        if (data.isSuccess) {
            setIsOpen(true)
        }
    }, [data])

    const clickHandler: MouseEventHandler<HTMLButtonElement> = (e) => {
        e.stopPropagation()
        addToCart(productId)
        console.log(`id:${productId} added to the buset`)
    }

    function modalClickHandler() {
        setIsOpen(false)
    }

    if (isIconOnly) {
        return (
            <ButtonNextUI isIconOnly onClick={clickHandler}>
                <PacketIcon />
            </ButtonNextUI>
        )
    }

    return (
        <>
            <Button onClick={clickHandler}>Добавить в корзину</Button>
            <Modal isOpen={isOpen} className='bg-slate-950'>
                <ModalHeader>Товар успешно добавлен в корзину!</ModalHeader>
                <Divider orientation='horizontal' />
                <ModalBody>
                    Отправленный запрос на добавление товара в корзину успешно
                    отработал, товар был добавлен в корзину
                </ModalBody>
                <Divider orientation='horizontal' />
                <ModalContent>
                    <ButtonNextUI color='primary' onPress={modalClickHandler}>
                        Закрыть
                    </ButtonNextUI>
                </ModalContent>
            </Modal>
        </>
    )
}

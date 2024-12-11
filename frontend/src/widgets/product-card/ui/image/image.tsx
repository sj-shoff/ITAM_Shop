import { Skeleton } from "@nextui-org/react"
import { useState } from "react"
import classes from "./image.module.scss"

type ImageProps = {
    path: string
}

export function Image({ path }: ImageProps) {
    const [isLoaded, setIsLoaded] = useState<boolean>(false)

    function onImageLoad() {
        setIsLoaded(!isLoaded)
    }

    return (
        <Skeleton className={classes.imageContainer} isLoaded={isLoaded}>
            <img
                className={classes.image}
                src={`data:image/jpeg;base64,${path}`}
                onLoad={onImageLoad}
                alt='product image'
                width='550'
                height='530'
                loading='lazy'
            />
        </Skeleton>
    )
}

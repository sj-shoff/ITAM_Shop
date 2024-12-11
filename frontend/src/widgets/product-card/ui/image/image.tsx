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
                className={`data:image/png;base64,${path}`}
                src={path}
                onLoad={onImageLoad}
                alt='product image'
                width='670'
                height='650'
                loading='lazy'
            />
        </Skeleton>
    )
}

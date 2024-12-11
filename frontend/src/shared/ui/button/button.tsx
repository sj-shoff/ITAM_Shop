import { HTMLAttributes, ReactNode } from "react"
import classes from "./button.module.scss"

type ButtonProps = HTMLAttributes<HTMLButtonElement> & {
    children?: ReactNode
    isIconOnly?: boolean
}

export function Button({
    children,
    isIconOnly = false,
    className = "",
    ...rest
}: ButtonProps) {
    return (
        <>
            {!isIconOnly ? (
                <button
                    type='button'
                    className={`${classes.button} ${className}`}
                    {...rest}
                >
                    {children}
                </button>
            ) : (
                <button
                    className={`${classes.iconButton} ${className}`}
                    {...rest}
                >
                    {children}
                </button>
            )}
        </>
    )
}

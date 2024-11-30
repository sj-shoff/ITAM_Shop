import { HTMLAttributes, HTMLProps, ReactNode } from "react"
import classes from "./button.module.scss"

type ButtonProps = HTMLProps<HTMLButtonElement> & {
    children?: ReactNode
}

export function Button({
    children,
    onClick,
    ...rest
}: ButtonProps & HTMLAttributes<HTMLButtonElement>) {
    return (
        <button className={classes.button} onClick={onClick} {...rest}>
            {children}
        </button>
    )
}

import { HTMLAttributes, HTMLProps, ReactNode } from "react"
import classes from "./button.module.scss"

type ButtonProps = HTMLProps<HTMLButtonElement> & {
    children?: ReactNode
}

export function Button(props: ButtonProps & HTMLAttributes<HTMLButtonElement>) {
    return (
        <button className={classes.button} onClick={props.onClick}>
            {props.children}
        </button>
    )
}

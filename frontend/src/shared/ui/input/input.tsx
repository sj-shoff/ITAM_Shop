import { HTMLProps } from "react"
import classes from "./input.module.scss"

type inputProps = HTMLProps<HTMLInputElement>

export function Input({ placeholder = "placeholder", ...rest }: inputProps) {
    return (
        <input
            className={classes.input}
            type='text'
            placeholder={placeholder}
            {...rest}
        />
    )
}

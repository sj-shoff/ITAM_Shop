import { HTMLAttributes, ReactNode } from "react"

type ButtonProps = {
    children: ReactNode
    onClick?: () => void
}

export function Button(props: ButtonProps & HTMLAttributes<HTMLButtonElement>) {
    return <button onClick={props.onClick}>{props.children}</button>
}

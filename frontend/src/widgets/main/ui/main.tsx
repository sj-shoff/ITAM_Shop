import classes from "./main.module.scss"

type Props = {
    children?: React.ReactNode
}

export function Main(props: Props) {
    return <main className={classes.main}>{props.children}</main>
}

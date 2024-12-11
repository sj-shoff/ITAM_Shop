import { Button } from "@shared/ui/button"
import { BasketIcon, PacketIcon, UserIcon } from "@shared/ui/icons"
import classes from "./buttons-group.module.scss"
import { Link } from "react-router-dom"

export function ButtonsGroup() {
    return (
        <ul className={classes.buttonsGroup}>
            <li>
                <Button className={classes.button} isIconOnly>
                    <BasketIcon />
                </Button>
            </li>
            <li>
                <Button className={classes.button} isIconOnly>
                    <UserIcon />
                </Button>
            </li>
            <li>
                <Link to='cart'>
                    <Button className={classes.button} isIconOnly>
                        <PacketIcon />
                    </Button>
                </Link>
            </li>
        </ul>
    )
}

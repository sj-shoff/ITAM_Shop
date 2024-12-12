import { Button } from "@shared/ui/button"
import { BasketIcon, PacketIcon, UserIcon } from "@shared/ui/icons"
import classes from "./buttons-group.module.scss"
import { Link } from "react-router-dom"

export function ButtonsGroup() {
    return (
        <ul className={classes.buttonsGroup}>
            <li>
                <Link className='contents' to='/fav_items'>
                    <Button className={classes.button} isIconOnly>
                        <BasketIcon />
                    </Button>
                </Link>
            </li>
            <li>
                <Link className='contents' to='/catalog'>
                    <Button className={classes.button} isIconOnly>
                        <UserIcon />
                    </Button>
                </Link>
            </li>
            <li>
                <Link className='contents' to='/cart'>
                    <Button className={classes.button} isIconOnly>
                        <PacketIcon />
                    </Button>
                </Link>
            </li>
        </ul>
    )
}

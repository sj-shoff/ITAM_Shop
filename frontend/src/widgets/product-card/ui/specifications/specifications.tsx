import { Divider } from "@nextui-org/react"
import classes from "./specifications.module.scss"
import { Feature } from "@entities/product"

export function Specifications({ features }: { features: Feature[] }) {
    return (
        <>
            {features.map((el) => (
                <li
                    key={el.name_of_feature}
                    className={classes.specificationItem}
                >
                    <div className={classes.characteristic}>
                        <span className={classes.name}>
                            {el.name_of_feature}
                        </span>
                        <span className={classes.value}>
                            {el.value_for_feature}
                        </span>
                    </div>
                    <Divider orientation='horizontal' />
                </li>
            ))}
        </>
    )
}

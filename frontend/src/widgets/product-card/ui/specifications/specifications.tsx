import { Divider } from "@nextui-org/react"
import classes from "./specifications.module.scss"
import { Feature } from "@entities/product"

export function Specifications({ features }: { features: Feature[] | null }) {
    let id = 1

    if (!features) {
        return "Спецификации отсутствуют"
    }
    return (
        <>
            {features.map((el) => (
                <li key={id++} className={classes.specificationItem}>
                    <div className={classes.characteristic}>
                        <span className={classes.name}>
                            {el.name_of_feature}
                        </span>
                        <div className={classes.value}>
                            <span className={classes.textValue}>
                                {el.value_for_feature}
                            </span>
                            <span className={classes.measurement}>
                                {el.Unit_of_measurement}{" "}
                            </span>
                        </div>
                    </div>
                    <Divider orientation='horizontal' />
                </li>
            ))}
        </>
    )
}

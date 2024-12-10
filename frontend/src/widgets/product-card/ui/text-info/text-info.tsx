import {
    ProductCategory,
    ProductDescription,
    ProductName,
} from "@entities/product"
import classes from "./text-info.module.scss"

export function TextInfo({
    name,
    category,
    description,
}: {
    name: ProductName
    category: ProductCategory
    description: ProductDescription
}) {
    return (
        <div className={classes.textInfo}>
            <h1 className={classes.name}>{name}</h1>
            <p className={classes.category}>{category}</p>
            <p className={classes.description}>{description}</p>
        </div>
    )
}

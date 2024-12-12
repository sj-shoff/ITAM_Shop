import { z } from "zod"

// ------------CATALOG---------------------------------------
export type ProductId = number
export type ProductName = string
export type ProductPrice = number
export type ProductImage = string
export type ProductDescription = string
export type ProductCategory = string
export type ProductStockQuantity = number
export type ProductIsInCart = number
export type ProductIsInFav = number

export type Product = {
    product_id: ProductId
    product_name: ProductName
    product_price: ProductPrice
    product_image: ProductImage
    product_description: ProductDescription
    product_category: ProductCategory
    product_stock_quantity: ProductStockQuantity
    is_in_cart: ProductIsInCart
    is_in_fav: ProductIsInFav
}

export const productDefaultValue: Product = {
    product_id: 0,
    product_name: "",
    product_price: 0,
    product_image: "",
    product_description: "",
    product_category: "",
    product_stock_quantity: 0,
    is_in_cart: 0,
    is_in_fav: 0,
}

export const productDTOschema = z.object({
    product_id: z.number(),
    product_price: z.number(),
    product_name: z.string(),
    product_image: z.string(),
    product_description: z.string(),
    product_category: z.string(),
    product_stock_quantity: z.number(),
    is_in_cart: z.number(),
    is_in_fav: z.number(),
})

// ------------FILTER---------------------------------------
export type FilterCategory = ProductCategory
export type FilterMinPrice = number
export type FilterMaxPrice = number

export type FilterOptions = {
    category: FilterCategory
    minPrice: FilterMinPrice
    maxPrice: FilterMaxPrice
}

// ------------FEATURES---------------------------------------
export type FeatureType = string
export type FeatureName = string
export type FeatureUnit = string

export type Feature = {
    name_of_feature: FeatureType
    value_for_feature: FeatureName
    Unit_of_measurement: FeatureUnit
}

export const featureDefaultValue: Feature = {
    name_of_feature: "",
    value_for_feature: "",
    Unit_of_measurement: "",
}

export const featureDTOschema = z.object({
    name_of_feature: z.string(),
    value_for_feature: z.string(),
    Unit_of_measurement: z.string(),
})

// ------------REQUEST---------------------------------------
export type RequestType = {
    features: Feature[] | null
    product: Product
}

export const requestTypeDefaultVale: RequestType = {
    features: [featureDefaultValue],
    product: productDefaultValue,
}

export const requestDTOschema = z.strictObject({
    features: featureDTOschema.array().nullable(),
    product: productDTOschema,
})

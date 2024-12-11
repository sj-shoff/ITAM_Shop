export type ProductId = number
export type ProductName = string
export type ProductPrice = number
export type ProductImage = string
export type ProductDescription = string
export type ProductCategory = string
export type ProductSpecifications = string
export type ProductStockQuantity = number

export type Product = {
    product_id: ProductId
    product_name: ProductName
    product_price: ProductPrice
    product_image: ProductImage
    product_description: ProductDescription
    product_category: ProductCategory
    product_specifications: ProductSpecifications
    product_stock_quantity: ProductStockQuantity
}

export type FilterCategory = ProductCategory
export type FilterMinPrice = number
export type FilterMaxPrice = number

export type FilterOptions = {
    category: FilterCategory
    minPrice: FilterMinPrice
    maxPrice: FilterMaxPrice
}

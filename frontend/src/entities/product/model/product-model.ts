export type ProductId = number
export type ProductName = string
export type ProductPrice = number
export type ProductImageUrl = string
export type ProductDescription = string
export type ProductCategory = string
export type ProductSpecifications = object
export type ProductStockQuantity = number

export type Product = {
    product_id: number
    product_name: string
    product_price: number
    product_image: string
    product_description: string
    product_category: string
    product_specifications: object
    product_stock_quantity: number
}

package entity

import "gorm.io/gorm"

type User struct {
	ID          uint    `json:"user_id" gorm:"column:user_id"`
	Balance     float64 `json:"user_balance" gorm:"column:user_balance"`
	UserName    string  `json:"user_name" gorm:"column:user_name"`
	UserSurname string  `json:"user_surname" gorm:"column:user_surname"`
	Login       string  `json:"user_login" gorm:"column:user_login"`
	Email       string  `json:"user_email" gorm:"column:user_email"`
	Password    string  `json:"user_password" gorm:"column:user_password"`
	Admin       bool    `json:"user_admin_rights" gorm:"column:user_admin_rights"`
	Avatar      []byte  `json:"user_avatar" gorm:"column:user_avatar"`
	//
}

type FilterParams struct {
	Category string  `json:"category"`
	MinPrice float64 `json:"minPrice"`
	MaxPrice float64 `json:"maxPrice"`
	//
}

type Product struct {
	gorm.Model
	ProductID      uint    `json:"product_id"`
	Price          float64 `json:"product_price"`
	Name           string  `json:"product_name"`
	ImageURL       string  `json:"image_url"`
	Description    string  `json:"product_description"`
	Category       string  `json:"product_category"`
	Specifications string  `json:"product_specifications"`
	Quantity       int     `json:"product_quantity"`
	StockQuantity  int     `json:"product_stock_quantity"`
	//
}

type CartItem struct {
	ProductID_cart uint    `json:"product_id"`
	Quantity       int     `json:"quantity"`
	Product        Product `json:"product"`
	//
}

type Cart struct {
	Items []CartItem `json:"items"`
	//
}

type Favorite struct {
	gorm.Model
	UserID uint `gorm:"not null" json:"user_id"`
	ItemID uint `gorm:"not null" json:"item_id"`
	//
}

type Order struct {
	gorm.Model
	UserID uint   `gorm:"not null" json:"user_id"`
	Status string `gorm:"not null" json:"status"`
	//
}

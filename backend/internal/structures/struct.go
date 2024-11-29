package entity

import "gorm.io/gorm"

type User struct {
	ID          string `json:"user_id"`
	UserName    string `json:"user_name"`
	UserSurname string `json:"user_surname"`
	Login       string `json:"user_login"`
	Email       string `json:"user_email"`
	Password    string `json:"user_password"`
	Admin       bool   `json:"user_admin_rights"`
	Avatar      string `json:"avatar"`
	//

}
type Filter struct {
	//
}

type Product struct {
	gorm.Model
	ProductID      uint    `json:"product_id"`
	Name           string  `json:"product_name"`
	Price          float64 `json:"product_price"`
	ImageURL       string  `json:"image_url"`
	Description    string  `json:"product_description"`
	Category       string  `json:"category"`
	StockQuantity  int     `json:"stock_quantity"`
	Specifications string  `json:"specifications"`
}

type CartItem struct {
	ProductID_cart uint    `json:"product_id"`
	Quantity       int     `json:"quantity"`
	Product        Product `json:"product"`
}

type Cart struct {
	Items []CartItem `json:"items"`
}

type Favorite struct {
	gorm.Model
	UserID uint `gorm:"not null" json:"user_id"`
	ItemID uint `gorm:"not null" json:"item_id"`
}

type Order struct {
	gorm.Model
	UserID uint   `gorm:"not null" json:"user_id"`
	Status string `gorm:"not null" json:"status"`
}

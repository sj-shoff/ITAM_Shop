package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	//gorm.Model
	ID          uint    `json:"user_id" gorm:"column:user_id"`
	Balance     float64 `json:"user_balance" gorm:"column:user_balance"`
	UserName    string  `json:"user_name" gorm:"column:user_name"`
	UserSurname string  `json:"user_surname" gorm:"column:user_surname"`
	Login       string  `json:"user_login" gorm:"column:user_login"`
	Email       string  `json:"user_email" gorm:"column:user_email"`
	Password    string  `json:"user_password" gorm:"column:user_password"`
	Admin       bool    `json:"user_admin_rights" gorm:"column:user_admin_rights"`
	Avatar      uint    `json:"user_avatar" gorm:"column:user_avatar"`
	//
}

type FilterParams struct {
	Category string  `json:"category" gorm:"column:category"`
	MinPrice float64 `json:"minPrice" gorm:"column:min_price"`
	MaxPrice float64 `json:"maxPrice" gorm:"column:max_price"`
	//
}

type Product struct {
	//gorm.Model
	//Features       []Feature `json:"features"`
	ProductID   uint    `json:"product_id" gorm:"column:product_id"`
	Price       float64 `json:"product_price" gorm:"column:product_price"`
	Name        string  `json:"product_name" gorm:"column:product_name"`
	Image       []byte  `json:"product_image" gorm:"column:product_image"`
	Description string  `json:"product_description" gorm:"column:product_description"`
	Category    string  `json:"product_category" gorm:"column:product_category"`
	//	Specifications string  `json:"product_specifications" gorm:"column:product_specifications"`
	Quantity      int `json:"product_quantity" gorm:"column:product_quantity"`
	StockQuantity int `json:"product_stock_quantity" gorm:"column:product_stock_quantity"`
	Is_in_cart    int `json:"is_in_cart" gorm:"column:is_in_cart"`
	Is_in_fav     int `json:"is_in_fav" gorm:"column:is_in_fav"`

	//
}

type Images struct {
	gorm.Model
	ImageID   uint   `json:"image_id" gorm:"column:image_id"`
	ImageData []byte `json:"image_data" gorm:"column:image_data"`
	//
}

type Feature struct {
	Name                string `json:"name_of_feature" gorm:"column:name_of_feature"`
	Value               string `json:"value_for_feature" gorm:"column:value_for_feature"`
	Unit_of_measurement string `json:"Unit_of_measurement" gorm:"column:Unit_of_measurement"`
}

type CartItem struct {
	ProductID_cart uint    `json:"product_id" gorm:"column:product_id"`
	Quantity       int     `json:"quantity" gorm:"column:quantity"`
	Product        Product `json:"product" gorm:"foreignKey:ProductID_cart"`
	//
}

type Cart struct {
	Items []CartItem `json:"items" gorm:"column:items"`
	//
}

type Favorite struct {
	gorm.Model
	ProductID      uint    `json:"product_id" gorm:"column:product_id"`
	Price          float64 `json:"product_price" gorm:"column:product_price"`
	Name           string  `json:"product_name" gorm:"column:product_name"`
	Image          uint    `json:"product_image" gorm:"column:product_image"`
	Description    string  `json:"product_description" gorm:"column:product_description"`
	Category       string  `json:"product_category" gorm:"column:product_category"`
	Specifications string  `json:"product_specifications" gorm:"column:product_specifications"`
	Quantity       int     `json:"product_quantity" gorm:"column:product_quantity"`
	StockQuantity  int     `json:"product_stock_quantity" gorm:"column:product_stock_quantity"`
	IsFavorite     bool    `json:"is_favorite" gorm:"column:is_favorite"`
}

type Order struct {
	gorm.Model
	UserID uint   `gorm:"not null" json:"user_id"`
	Status string `gorm:"not null" json:"status"`
	//
}

type Sale struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primaryKey;column:id"`
	ProductID uint      `json:"product_id" gorm:"column:product_id"`
	Quantity  int       `json:"quantity" gorm:"column:quantity"`
	Price     float64   `json:"price" gorm:"column:price"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

package entity

type User struct {
	ID          string `json:"user_id"`
	UserName    string `json:"user_name"`
	UserSurname string `json:"user_surname"`
	Login       string `json:"user_login"`
	Email       string `json:"user_email"`
	Password    string `json:"user_password"`
	Admin       bool   `json:"user_admin_rights"`
	//

}

type Product struct {
	ProductID   uint    `json:"product_id"`
	Name        string  `json:"product_name"`
	Price       float64 `json:"product_price"`
	Description string  `json:"product_discription"`
	//
}

type Filter struct {
	//
}
type CartItem struct {
	ProductID_cart uint    `json:"product_id_cart"`
	Quantity       int     `json:"quantity"`
	Product        Product `json:"product_in_cart"`
}

type Cart struct {
	ID    uint       `json:"id"`
	Items []CartItem `json:"items"`
}

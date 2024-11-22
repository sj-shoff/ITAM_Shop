package entity

type User struct {
	ID          string
	UserName    string
	UserSurname string
	Login       string
	Password    string
	Admin       bool
	//

}
type Product struct {
	ID          string
	Name        string
	Cost        uint
	Description string
	//
}

type Filter struct {
	//
}

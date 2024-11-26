package main

import (
	"log"
	storage "myapp/internal"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func init() {
	var err error
	var db *gorm.DB
	dsn := "postgresql://user:password@localhost/database_name?sslmode=disable" // Настроим потом
	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		panic("Ошибка подключения к базе данных")
	}
	db.Debug()
}
func main() {

	r := gin.Default()

	r.GET("/front_page", storage.ShowHomePage)
	r.GET("/catalog")

	r.GET("/register", storage.ShowRegistrationForm)
	r.POST("/register")
	r.GET("/login", storage.ShowLoginForm)
	r.POST("/login")
	r.GET("/login/:id/acc")

	r.POST("/catalog/filter")
	r.GET("/catalog/fav_items")
	r.POST("/catalog/fav_items/:id")

	r.POST("/cart/item/:id")
	r.DELETE("/cart/item/:id")
	r.GET("/cart")

	r.GET("/analytics")
	r.GET("/admin_panel")

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}
	log.Println("Сервер запущен на http://localhost:8080")

}

package main

import (
	"log"
	storage "myapp/internal"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	var err error
	dsn := "admin_for_itam_store:your_password@tcp(147.45.163.58:3306)/itam_store?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных:", err)
	}
	db.Debug()
}
func main() {

	initDB()

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

	r.POST("/cart/item/:id", storage.AddToCart(db))
	r.DELETE("/cart/item/:id", storage.RemoveFromCart(db))
	r.GET("/cart", storage.ShowCart)

	r.GET("/analytics")
	r.GET("/admin_panel")

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}
	log.Println("Сервер запущен на http://localhost:8080")

}

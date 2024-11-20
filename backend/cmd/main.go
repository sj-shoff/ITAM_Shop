package main

import (
	"log"

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
	// , _ - под функции из server
	r.GET("/")         // главная страница
	r.GET("/register") // регистрация
	r.POST("/register")
	r.GET("/login") // авторизация
	r.POST("/login")

	r.GET("/catalog")       // каталог
	r.POST("/add_item/:id") // добавление айтема в корзину

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}
	log.Println("Сервер запущен на http://localhost:8080")

}

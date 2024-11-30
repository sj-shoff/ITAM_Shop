package main

import (
	"log"
//	"fmt"
	storage "myapp/internal"

	"github.com/gin-gonic/gin"
//	"github.com/jinzhu/gorm"
//"database/sql"

//    "github.com/go-sql-driver/mysql"
)


func init() {

}
func main() {

	r := gin.Default()
	//r.Static("/templates", "./f")
	r.LoadHTMLFiles("../f/index.html")
	r.LoadHTMLFiles("../f/add_new_product.html")
	r.GET("/", storage.ShowHomePage)




	r.GET("/catalog")

	r.GET("/register", storage.ShowRegistrationForm)
	r.POST("/register")
	r.GET("/login", storage.ShowLoginForm)
	r.POST("/login")
	r.GET("/login/:id/acc")

	r.POST("/catalog/filter")
	r.GET("/catalog/fav_items")
	r.POST("/catalog/fav_items/:id")

	r.POST("/add_item", storage.CreateNewProduct)
	r.DELETE("/cart/:id")

	r.GET("/analytics")
	r.GET("/admin_panel")

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}
	log.Println("Сервер запущен на http://localhost:8080")

}

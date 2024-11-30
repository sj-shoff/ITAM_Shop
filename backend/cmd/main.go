package main

import (
	"log"
//	"fmt"
	storage "myapp/internal"
	catalog "myapp/internal/catalog"
	config "myapp/internal/data_base"
	controllers "myapp/internal/personal_account/controllers"

	"github.com/gin-gonic/gin"

//	"github.com/jinzhu/gorm"
//"database/sql"

//    "github.com/go-sql-driver/mysql"
)


func init() {

}
=======

func main() {

	config.InitDB()
	r := gin.Default()

	//r.Static("/templates", "./f")
	r.LoadHTMLFiles("../f/index.html")
	r.LoadHTMLFiles("../f/add_new_product.html")
	r.GET("/", storage.ShowHomePage)




	r.GET("/catalog")


	//регистрация-авторизация
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

	// логика лк
	r.GET("/favorites/:id", controllers.GetFavorites)
	r.PUT("/users/:id/avatar", controllers.UpdateAvatar)
	r.PUT("/users/:id/password", controllers.UpdatePassword)
	r.POST("/logout", controllers.Logout)
	r.GET("/orders/:id", controllers.GetOrders)
	r.GET("/orders/:id/status", controllers.GetOrderStatus)
	r.GET("/favorites/:id", controllers.GetFavorites)

	//логика каталога
	r.GET("/products", catalog.GetProducts)
	r.GET("/products/:id", catalog.GetProduct)
	r.POST("/products/filter", catalog.ProductFilter)
	r.POST("/cart/item/:id", catalog.AddToCart)
	r.DELETE("/cart/item/:id", catalog.RemoveFromCart)
	r.GET("/cart", storage.ShowCart)
	r.POST("/favorites", catalog.AddToFavorites)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}
	log.Println("Сервер запущен на http://localhost:8089")

}

package main

import (
	"log"
	storage "myapp/internal"
	config "myapp/internal/data_base"
	controllers "myapp/internal/personal_account/controllers"
	entity "myapp/internal/structures"

	"github.com/gin-gonic/gin"
)

func main() {

	config.InitDB()
	config.DB.AutoMigrate(&entity.User{}, &entity.Order{}, &entity.Favorite{})

	r := gin.Default()

	r.GET("/front_page", storage.ShowHomePage)
	r.GET("/catalog")

	r.GET("/register", storage.ShowRegistrationForm)
	r.POST("/register")
	r.GET("/login", storage.ShowLoginForm)
	r.POST("/login")
	r.GET("/login/:id/acc")

	r.POST("/catalog/filter")

	r.GET("/favorites/:id", controllers.GetFavorites)

	r.POST("/cart/item/:id", storage.AddToCart(config.DB, entity.CartItem{}))
	r.DELETE("/cart/item/:id", storage.RemoveFromCart(config.DB))
	r.GET("/cart", storage.ShowCart)

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

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}
	log.Println("Сервер запущен на http://localhost:8080")

}

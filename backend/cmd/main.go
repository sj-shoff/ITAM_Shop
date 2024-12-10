package main

import (
	"log"
	storage "myapp/internal"
	config "myapp/internal/data_base"
	"myapp/internal/personal_account/controllers"
	"myapp/internal/register"
	entity "myapp/internal/structures"
	"myapp/transactions"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	config.InitDB()
	config.DB.AutoMigrate(&entity.User{}, &entity.Order{}, &entity.Favorite{})

	r := gin.Default()

	// Health
	r.GET("/health", func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusOK)
	})

	r.GET("/front_page", storage.ShowHomePage)
	r.GET("/catalog")

	// Sessions
	store := cookie.NewStore([]byte("secret-key"))
	r.Use(sessions.Sessions("mysession", store))

	// Register, Login, Sessions
	register.InitRegister(config.DB, r)

	// Personal account
	controllers.InitPersonalAccount(config.DB, r)

	// Transactions
	transactions.InitTransaction(config.DB, r)

	//

	//r.GET("/register", storage.ShowRegistrationForm)
	//r.GET("/login", storage.ShowLoginForm)
	//r.GET("/login/:id/acc")

	r.POST("/catalog/filter")
	r.GET("/catalog/fav_items")
	r.POST("/catalog/fav_items/:id")

	//r.POST("/cart/item/:id", storage.AddToCart)
	//r.DELETE("/cart/item/:id", storage.RemoveFromCart)
	//r.GET("/cart", storage.ShowCart)

	r.GET("/analytics")
	r.GET("/admin_panel")

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}
	log.Println("Сервер запущен на http://localhost:8080")

}

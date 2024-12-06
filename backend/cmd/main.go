package main

import (
	"log"
	admin "myapp/internal/admin_panel"
	catalog "myapp/internal/catalog"
	config "myapp/internal/data_base"
	"myapp/internal/personal_account/controllers"
	"myapp/internal/register"
	entity "myapp/internal/structures"
	transactions "myapp/internal/transactions"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	// Init DB
	config.InitDB()
	config.DB.AutoMigrate(&entity.User{}, &entity.Order{}, &entity.Favorite{})

	r := gin.Default()

	// Health
	r.GET("/health", func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusOK)
	})

	// Sessions
	store := cookie.NewStore([]byte("secret-key"))
	r.Use(sessions.Sessions("mysession", store))

	// Register, Login, Sessions
	register.InitRegister(config.DB, r)

	// Personal account
	controllers.InitPersonalAccount(config.DB, r)

	// Transactions
	transactions.InitTransaction(config.DB, r)

	// Catalog
	catalog.InitCatalog(config.DB, r)

	r.GET("/front_page") // пока не сделано
	r.Use(func(c *gin.Context) {
		c.Header("Content-Security-Policy", "default-src 'self' http://localhost:3000;")
	})
	r.GET("/analytics")
	r.GET("/admin_panel", admin.AdminPanelHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}
	log.Println("Сервер запущен на http://localhost:8080")

}

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

	"github.com/gin-contrib/cors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	// Init DB
	config.InitDB()
	config.DB.AutoMigrate(&entity.User{}, &entity.Order{}, &entity.Favorite{})

	r := gin.Default()
	r.Use(cors.Default())
	/*
		r.Use(func(c *gin.Context) {
			c.Header("Access-Control-Allow-Origin", "*")                   // Разрешаем все домены
			c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE OPTIONS") // Разрешаем методы
			c.Header("Access-Control-Allow-Headers", "Content-Type")       // Разрешаем заголовки

			// Если это preflight-запрос, просто возвращаем 200 OK
			if c.Request.Method == http.MethodOptions {
				c.AbortWithStatus(http.StatusOK)
				return
			}

			c.Next()
		})
	*/
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
	admin.InitAdminPanel(config.DB, r)
	// Catalog
	catalog.InitCatalog(config.DB, r)

	r.GET("/front_page") // пока не сделано
	r.Use(func(c *gin.Context) {
		c.Header("Content-Security-Policy", "default-src 'self' http://localhost:8080;")
	})
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

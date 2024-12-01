package controllers

import (
	config "myapp/internal/data_base"
	entity "myapp/internal/structures"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitPersonalAccount(db1 *gorm.DB, s *gin.Engine) {
	db = db1
	s.POST("/updateavatar", UpdateAvatar)
	s.POST("/updatename", UpdateName())
	s.POST("/updatesurname", UpdateSurname())
	s.POST("/updatepassword", UpdatePassword)
	s.POST("/logout", Logout)
}

func UpdateAvatar(c *gin.Context) {
	sessions := sessions.Default(c)
	id := sessions.Get("id")
	if id == nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "User not logged"})
		return
	}

	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}

	if err := config.DB.Model(&entity.User{}).Where("user_id = ?", id).Update("user_avatar", user.Avatar).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Сообщение": "Аватарка обновлена"})
}

func UpdatePassword(c *gin.Context) {
	var request struct {
		NewPassword string `json:"user_password"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request"})
		return
	}

	sessions := sessions.Default(c)
	login := sessions.Get("login")

	if err := db.Model(&entity.User{}).Where("user_login = ?", login).Update("user_password", request.NewPassword).Error; err != nil {
		c.JSON(400, gin.H{"message": "Error"})
		//log.Print("Database error: ", err)
		return
	}

	c.JSON(200, gin.H{"message": "Password updated"})
}

func UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			Name string `json:"user_name"`
		}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid request"})
			return
		}

		sessions := sessions.Default(ctx)
		login := sessions.Get("login")

		if err := db.Model(&entity.User{}).Where("user_login = ?", login).Update("user_name", request.Name).Error; err != nil {
			ctx.JSON(400, gin.H{"message": "Error"})
			return
		}

		ctx.JSON(200, gin.H{"message": "Password updated"})
	}
}

func UpdateSurname() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			Name string `json:"user_surname"`
		}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid request"})
			return
		}

		sessions := sessions.Default(ctx)
		login := sessions.Get("login")

		if err := db.Model(&entity.User{}).Where("user_login = ?", login).Update("user_surname", request.Name).Error; err != nil {
			ctx.JSON(400, gin.H{"message": "Error"})
			return
		}

		ctx.JSON(200, gin.H{"message": "Password updated"})
	}
}

func Logout(c *gin.Context) {
	sessions := sessions.Default(c)

	sessions.Delete("id")
	sessions.Delete("login")
	sessions.Delete("code")

	if err := sessions.Save(); err != nil {
		c.JSON(500, gin.H{"message": "Error saving session"})
		return
	}

	c.JSON(200, gin.H{"message": "Successfully logged out"})
}

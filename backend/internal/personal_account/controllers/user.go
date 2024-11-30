package controllers

import (
	config "myapp/internal/data_base"
	entity "myapp/internal/structures"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Недействительный ID пользователя"})
		return
	}
	var user entity.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Ошибка": "Пользователь не найден"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateAvatar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}

	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})

		return
	}

	if err := config.DB.Model(&entity.User{}).Where("id = ?", id).Update("avatar", user.Avatar).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Сообщение": "Аватарка обновлена"})
}

func UpdatePassword(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Недействительный ID пользователя"})
		return
	}

	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&entity.User{}).Where("id = ?", id).Update("password", user.Password).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated"})
}
func Logout(c *gin.Context) {
	// я блять не ебу что тут писать
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}

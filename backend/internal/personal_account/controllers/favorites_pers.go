package controllers

import (
	config "myapp/internal/data_base"
	entity "myapp/internal/structures"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetFavorites(c *gin.Context) {
	sessions := sessions.Default(c)
	id := sessions.Get("id")
	if id == nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Недействительный ID пользователя"})
		return
	}

	var favorites []entity.Favorite
	if err := config.DB.Where("user_id = ?", id).Find(&favorites).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
		return
	}

	c.JSON(http.StatusOK, favorites)
}

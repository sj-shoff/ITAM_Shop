package controllers

import (
	config "myapp/internal/data_base"
	entity "myapp/internal/structures"
	"net/http"
	"strconv"

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

func GetOrders(c *gin.Context) {
	sessions := sessions.Default(c)
	id := sessions.Get("id")
	if id == nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Недействительный ID пользователя"})
		return
	}

	var orders []entity.Order
	if err := config.DB.Where("user_id = ?", id).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func GetOrderStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Недействительный ID заказа"})
		return
	}

	var order entity.Order
	if err := config.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Ошибка": "Заказ не найден"})
		return
	}

	c.JSON(http.StatusOK, order)
}

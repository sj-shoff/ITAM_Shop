package controllers

import (
	config "myapp/internal/data_base"
	entity "myapp/internal/structures"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Недействительный ID пользователя"})

		return
	}

	var cart []entity.Cart
	if err := config.DB.Where("user_id = ?", id).Find(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})

		return
	}

	c.JSON(http.StatusOK, cart)
}

func AddToCart(c *gin.Context) {
	var cartItem entity.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})

		return
	}

	if err := config.DB.Create(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"Сообщение": "Продукт добавлен в корзину"})
}

func RemoveFromCart(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Недействительный ID продукта"})

		return
	}

	if err := config.DB.Where("id= ?", id).Delete(&entity.CartItem{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"Сообщение": "Продукт удален из карзины"})
}

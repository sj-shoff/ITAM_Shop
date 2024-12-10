package server

import (
	"database/sql"
	"strconv"

	//"encoding/json"
	entity "myapp/internal/structures"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var cart entity.Cart

func ShowHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func ShowRegistrationForm(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func ShowLoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func ShowCart(c *gin.Context) {
	c.HTML(http.StatusOK, "cart.html", nil)
}

func Add(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product entity.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Преобразуем данные в JSON
		// productJSON, err := json.Marshal(product)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
		// 	return
		// }

		// Сохраняем данные в базу данных
		// query := `
		// 	INSERT INTO products (name, start_date, end_date, update_date, version_description, series_prefix, series_postfix, number_prefix, number_postfix, numerator, custom_number_method, individual_parameters, cost_formula)
		// 	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		// `
		// _, err = db.Exec(query, product.Name, product.StartDate, product.EndDate, product.UpdateDate, product.VersionDescription, product.MandatoryParams, productJSON, product.CostFormula)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data"})
		// 	return
		// }

		c.JSON(http.StatusOK, gin.H{"message": "Данные успешно отправлены"})
	}
}

func AddToCart(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item entity.CartItem
		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
			return
		}
		var product entity.Product
		if err := db.First(&product, item.ProductID_cart).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"Ошибка": "Продукт не найден"})
			return
		}
		item.Product = product
		cart.Items = append(cart.Items, item)

		c.JSON(http.StatusOK, cart)
	}

}

func RemoveFromCart(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		productID_fromURL := c.Param("id")

		productID, err := strconv.ParseUint(productID_fromURL, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Неправильный ID продукта"})
		}
		for i, cartItem := range cart.Items {
			if cartItem.ProductID_cart == uint(productID) {
				cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
				c.JSON(http.StatusOK, cart)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"Ошибка": "Продукт не найден в корзине"})
	}
}

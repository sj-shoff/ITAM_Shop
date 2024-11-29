package server

import (
	entity "myapp/internal/structures"
	//config "myapp/internal/data_base"
	"strconv"
	//"encoding/json"

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

func AddToCart(db *gorm.DB, item entity.CartItem) gin.HandlerFunc {
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

// func AddToCart(db *gorm.DB, jsonFilePath string) error {
//     // Чтение JSON-файла
//     jsonFile, err := ioutil.ReadFile(jsonFilePath)
//     if err != nil {
//         return fmt.Errorf("Failed to read JSON file: %v", err)
//     }

//     // Десериализация JSON-файла в структуру CartItem
//     var item entity.CartItem
//     if err := json.Unmarshal(jsonFile, &item); err != nil {
//         return fmt.Errorf("Failed to unmarshal JSON: %v", err)
//     }

//     // Поиск продукта в базе данных
//     var product entity.Product
//     if err := db.First(&product, item.ProductID_cart).Error; err != nil {
//         return fmt.Errorf("Product not found: %v", err)
//     }

//     // Добавление продукта в корзину
//     item.Product = product
//     cart.Items = append(cart.Items, item)
//     return nil
// }

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

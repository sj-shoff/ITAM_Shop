package server

import (
	"database/sql"
	//"encoding/json"
	entity "myapp/internal/structures"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name, Surname, Id, Email, Password string
	Is_admin                           bool
}

func ShowHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func ShowRegistrationForm(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func ShowLoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func AddItemToCart(db *sql.DB) gin.HandlerFunc {
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

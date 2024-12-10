package server

import (

	//	"log"
	"fmt"
	//"encoding/json"

	entity "myapp/internal/structures"

	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var adress_data_base_test = "admin_for_itam_store:your_password@tcp(147.45.163.58:3306)/itam_store"

var cart entity.Cart

func ShowHomePage(c *gin.Context) {
	fmt.Println("good")
	var err error

	db, err := sql.Open("mysql", "admin_for_itam_store:your_password@tcp(147.45.163.58:3306)/itam_store")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Printf("Подключено")
	//Установка данных
	//insert, err := db.Query(fmt.Sprintf("INSERT INTO test.articles (`title`, `anons`, `full_text`) VALUES ('%s', '%s', '%s')", title, anons, full_text))
	// var zapros = fmt.Sprintf("SELEC T* FROM `users`")
	// _,err = db.Query(zapros)
	//fmt.Println(zapros)

	c.HTML(200, "index.html", gin.H{
		"title": "Main website", //IGNORE THIS
	})

}

func ShowRegistrationForm(c *gin.Context) {

	c.HTML(http.StatusOK, "register.html", nil)
}

func ShowLoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func AddToCart(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func CreateNewProduct(c *gin.Context) {
	fmt.Println("Strart")
	var newItem entity.Product
	if err := c.ShouldBindJSON(&newItem); err == nil {
		// Здесь .вы можете добавить логику для обработки нового элемента
		c.JSON(http.StatusCreated, newItem)
	} else {
		// Если произошла ошибка, возвращаем статус 400 с сообщением об ошибке
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(newItem.Name)
	fmt.Println(newItem.Price)
	fmt.Println(newItem.Description)
	db, err := sql.Open("mysql", adress_data_base_test)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into itam_store.products_in_store ( `name`, `price`, `description`, `quantity`) values (?, ?, ?, ?)", newItem.Name, newItem.Price, newItem.Description, newItem.Quantity)

	fmt.Println(result)
	if err != nil {
		fmt.Println(err)
	}

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

func sk(db *gorm.DB) gin.HandlerFunc {
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

func ShowFavourites(c *gin.Context) {
	c.HTML(http.StatusOK, "favourites.html", nil)
}

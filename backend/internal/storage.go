package server

import (


//	"strconv"
//	"log"
	"fmt"
	//"encoding/json"

	entity "myapp/internal/structures"

	"net/http"
	"database/sql"
  _ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
)


func GetAuthUser(c *gin.Context) entity.User {
	//sessions := sessions.Default(ctx)
	login := "Mover-R" //sessions.Get("login")
	var adress_data_base_test = "admin_for_itam_store:your_password@tcp(147.45.163.58:3306)/itam_store"

	db, err := sql.Open("mysql", adress_data_base_test)
	if err != nil{
		panic(err)
	}

	defer db.Close()
	var zapros = fmt.Sprintf("SELECT user_id, user_name FROM `users` WHERE user_login='%s'", login)
	res,err := db.Query(zapros)
	fmt.Println(zapros)
	fmt.Println(res)

	var AuthUser entity.User
	for res.Next(){
		err = res.Scan(&AuthUser.ID, &AuthUser.UserName)
		if err != nil{
			panic(err)
		}
	}
	return AuthUser
}
//var cart entity.Cart



func ShowHomePage(c *gin.Context) {

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



func ShowFavourites(c *gin.Context) {
	c.HTML(http.StatusOK, "favourites.html", nil)
}

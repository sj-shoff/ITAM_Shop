package controllers

import (
	config "myapp/internal/data_base"
	server "myapp/internal"
	entity "myapp/internal/structures"
	"net/http"
	"strconv"
	"fmt"
	"github.com/gin-gonic/gin"

	"database/sql"
  _ "github.com/go-sql-driver/mysql"
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
	fmt.Println("ADD TO CART")
	id_product, err := strconv.Atoi(c.Param("id"))
	fmt.Println("Strart")

	var adress_data_base_test = "admin_for_itam_store:your_password@tcp(147.45.163.58:3306)/itam_store"

	db, err := sql.Open("mysql", adress_data_base_test)
  if err != nil{
    panic(err)
  }
  defer db.Close()
	AuthUser := server.GetAuthUser(c)
	result, err := db.Exec("insert into itam_store.users_carts ( `user_id`, `products_id`, `quantity`) values (?, ?,?)",AuthUser.ID ,id_product ,1)

	fmt.Println(result)
	if(err != nil){
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"Сообщение": "Продукт добавлен в корзину"})
}

func RemoveFromCart(c *gin.Context) {
	fmt.Println("ADD TO CART")
	id_product, err := strconv.Atoi(c.Param("id"))
	fmt.Println("Strart")


	var adress_data_base_test = "admin_for_itam_store:your_password@tcp(147.45.163.58:3306)/itam_store"

	db, err := sql.Open("mysql", adress_data_base_test)
  if err != nil{
    panic(err)
  }
  defer db.Close()
	AuthUser := server.GetAuthUser(c)
	result, err := db.Exec("DELETE FROM users_carts WHERE products_id = ?", products_id)

	fmt.Println(result)
	if(err != nil){
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{"Сообщение": "Продукт удален из корзины"})
}

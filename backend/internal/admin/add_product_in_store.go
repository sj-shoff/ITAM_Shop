
package admin

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
	var adress_data_base_test = "admin_for_itam_store:your_password@tcp(147.45.163.58:3306)/itam_store"

	db, err := sql.Open("mysql", adress_data_base_test)
  if err != nil{
    panic(err)
  }
  defer db.Close()

	result, err := db.Exec("insert into itam_store.products_in_store ( `name`, `price`, `description`, `quantity`, `img`) values (?, ?,?, ?, ?)",newItem.Name ,newItem.Price ,newItem.Description,newItem.Quantity, newItem.ImageURL)

	fmt.Println(result)
	if(err != nil){
		fmt.Println(err)
	}


}

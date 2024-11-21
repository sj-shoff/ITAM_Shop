// You can edit this code!
// Click here and start typing.
package main

import ("fmt"
        "github.com/gin-gonic/gin"
		//	"database/sql"
		//_ "github.com/go-sql-driver/mysql"
		//  "github.com/gin-contrib/sessions"

			)

func main() {
	fmt.Println("Hello")

	r.Static("/static", "./frontend/static")
	r.Static("/templates", "./frontend/templates")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

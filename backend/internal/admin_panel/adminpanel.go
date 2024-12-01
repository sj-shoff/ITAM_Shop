package adminpanel

import (
	"log"
	entity "myapp/internal/structures"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitAdminPanel(db1 *gorm.DB, s *gin.Engine) {
	db = db1
	s.POST("/createnewprod", CreateProduct())
	s.POST("/editproduct/:id", EditProduct())
	s.POST("/deleteproduct/:id", DeleteProduct())
}

func CreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var p entity.Product

		if err := ctx.ShouldBindJSON(&p); err != nil {
			ctx.JSON(204, gin.H{"message": "Bad data for edit"})
			return
		}

		res := db.Create(&p)

		if res.Error != nil {
			ctx.JSON(500, gin.H{"message": "Failed create product"})
			return
		}

		ctx.JSON(200, gin.H{"message": "Product succesfully created"})
	}
}

func EditProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var p entity.Product

		if err := ctx.ShouldBindJSON(&p); err != nil {
			ctx.JSON(204, gin.H{"message": "Bad data for edit"})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Print(err)
			ctx.JSON(404, gin.H{"message": "No such product"})
			return
		}

		query := "product_name, product_price, product_description, product_category, product_specifications, product_quantity, product_stock_quantity"
		if err := db.Model(&entity.Product{}).Where("product_id = ?", id).Update(query, p).Error; err != nil {
			ctx.JSON(400, gin.H{"message": "Error"})
			log.Print("Database error: ", err)
			return
		}

		ctx.JSON(200, gin.H{"message": "Succesfully edited product"})
	}
}

func DeleteProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Print(err)
			ctx.JSON(404, gin.H{"message": "No such product"})
			return
		}

		var prod entity.Product

		res := db.Delete(&prod, id)
		if res.Error != nil {
			ctx.JSON(500, gin.H{"message": "failed to delete product"})
			return
		}

		ctx.JSON(200, gin.H{"message": "Succesfully deleted"})
	}
}

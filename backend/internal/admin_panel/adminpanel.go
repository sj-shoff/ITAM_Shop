package adminpanel

import (
	"fmt"
	"log"
	entity "myapp/internal/structures"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/gorm"
)

var (
	db *gorm.DB

	productsCreated = promauto.NewCounter(prometheus.CounterOpts{
		Name: "products_created_total",
		Help: "The total number of created products",
	})

	productsEdited = promauto.NewCounter(prometheus.CounterOpts{
		Name: "products_edited_total",
		Help: "The total number of edited products",
	})

	productsDeleted = promauto.NewCounter(prometheus.CounterOpts{
		Name: "products_deleted_total",
		Help: "The total number of deleted products",
	})

	salesTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "sales_total",
		Help: "The total number of sales",
	})

	siteVisitsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "site_visits_total",
		Help: "The total number of site visits",
	})
)

func AdminPanelHandler(c *gin.Context) {
	htmlFile, err := os.Open("/home/sj_shoff/ITAM_Shop/backend/f/admin_panel.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error reading HTML fileA")
		return
	}
	defer htmlFile.Close()

	htmlContent, err := os.ReadFile("/home/sj_shoff/ITAM_Shop/backend/f/admin_panel.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error reading HTML file")
		return
	}

	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, string(htmlContent))
}

func InitAdminPanel(db1 *gorm.DB, s *gin.Engine) {
	db = db1
	s.POST("/add_features_to_item/:id_item/:id_features", AddFeaturesToItem())

	s.POST("/createnewprod", CreateProduct())
	s.POST("/editproduct/:id", EditProduct())
	s.POST("/deleteproduct/:id", DeleteProduct())
	s.POST("/updateimageforproduct/:id", UpdateImageForProduct())

	s.GET("/metrics", gin.WrapH(promhttp.Handler()))
	s.GET("/", TrackSiteVisit())
	s.POST("/sale", TrackSale())
}

type Temp struct {
	Message string `json:"message" gorm:"column:message"`
}

// tracking
func AddFeaturesToItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id_item, err := strconv.Atoi(ctx.Param("id_item"))
		if err != nil {
			ctx.JSON(400, gin.H{"message": "Bad request"})
			return
		}
		id_features, err := strconv.Atoi(ctx.Param("id_features"))
		var value_struct Temp
		if err := ctx.ShouldBindJSON(&value_struct); err != nil {
			fmt.Println("Error")
			fmt.Println(err)
			ctx.JSON(204, gin.H{"message": "Bad data for edit"})
			return
		}
		result := db.Exec("insert into itam_store.added_features (id_item	, value, id_feature) values (?, ?,?)", id_item, value_struct.Message, id_features)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)

		ctx.JSON(200, gin.H{"message": "features added"})
	}
}

func TrackSiteVisit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		siteVisitsTotal.Inc()
		ctx.JSON(200, gin.H{"message": "Welcome to the site"})
	}
}

func TrackSale() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sale entity.Sale

		if err := ctx.ShouldBindJSON(&sale); err != nil {
			ctx.JSON(204, gin.H{"message": "Bad data for sale"})
			return
		}

		salesTotal.Inc()
		ctx.JSON(200, gin.H{"message": "Sale tracked"})
	}
}

// admin functions
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

func DeleteProductImage(id int) bool {
	var prod entity.Product

	if err := db.Model(&entity.Product{}).Where("product_id = ?", id).First(&prod).Error; err != nil {
		return false
	}

	ImageID := prod.Image
	query := "DELETE FROM images WHERE image_id = ?"
	if err := db.Exec(query, ImageID); err != nil {
		return false
	}

	return true
}

func UpdateImageForProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Print(err)
			ctx.JSON(404, gin.H{"message": "No such product"})
			return
		}

		var newImage entity.Images
		if err := ctx.ShouldBindJSON(&newImage); err != nil {
			ctx.JSON(400, gin.H{"message": "Bad request"})
			return
		}

		result := db.Create(&newImage)

		if result.Error != nil {
			ctx.JSON(400, gin.H{"message": "Ошибка"})
			return
		}

		if !DeleteProductImage(id) {
			ctx.JSON(400, gin.H{"message": "Bad request"})
			return
		}

		if err := db.Model(&entity.Product{}).Where("product_id = ?", id).Update("product_image = ?", newImage.ImageID).Error; err != nil {
			ctx.JSON(400, gin.H{"message": "Error"})
			return
		}

		ctx.JSON(200, gin.H{"message": "Product image updated"})
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

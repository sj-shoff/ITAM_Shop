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

	s.POST("/editproductname/:id", EditProductName())
	s.POST("/editproductdescription/:id", EditProductDescription())
	s.POST("/editproductprice/:id", EditProductPrice())
	s.POST("/editproductcategory/:id", EditProductCategory())
	s.POST("/editproductquantity/:id", EditProductQuantity())
	s.POST("/editproductstockquantity/:id", EditProductStockQuantity())

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
		id_item, err := strconv.Atoi(ctx.Param("id_item")) //Предмет к которому добавляем фичу
		if err != nil {
			ctx.JSON(400, gin.H{"message": "Bad request"})
			return
		}
		id_features, err := strconv.Atoi(ctx.Param("id_features"))//id фичи которую хотим добавить
		var value_struct Temp
		if err := ctx.ShouldBindJSON(&value_struct); err != nil { // получаем значения для фичи которую добавляем
			fmt.Println("Error")
			fmt.Println(err)
			ctx.JSON(204, gin.H{"message": "Bad data for edit"})
			return
		}
		result := db.Exec("insert into itam_store.added_features (id_item, value, id_feature) values (?, ?,?)", id_item, value_struct.Message, id_features)
		if err != nil { // Отправляем сведенья в бд
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

		if err := db.Model(&entity.Product{}).Where("product_id = ?", id).Update("product_image = ?", newImage.ImageData).Error; err != nil {
			ctx.JSON(400, gin.H{"message": "Error"})
			return
		}

		ctx.JSON(200, gin.H{"message": "Product image updated"})
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

func EditProductName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			Name string `json:"product_name"`
		}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid request"})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Print(err)
			ctx.JSON(404, gin.H{"message": "No such product"})
			return
		}

		if err := db.Model(&entity.Product{}).Where("product_id = ?", id).Update("product_name", request.Name).Error; err != nil {
			ctx.JSON(400, gin.H{"message": "Error"})
			return
		}

		ctx.JSON(200, gin.H{"message": "Product Name updated"})
	}
}

func EditProductPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			Price int `json:"product_price"`
		}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid request"})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Print(err)
			ctx.JSON(404, gin.H{"message": "No such product"})
			return
		}

		if err := db.Model(&entity.Product{}).Where("product_id = ?", id).Update("product_price", request.Price).Error; err != nil {
			ctx.JSON(400, gin.H{"message": "Error"})
			return
		}

		ctx.JSON(200, gin.H{"message": "Product Price updated"})
	}
}

func EditProductDescription() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			D string `json:"product_description"`
		}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid request"})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Print(err)
			ctx.JSON(404, gin.H{"message": "No such product"})
			return
		}

		if err := db.Model(&entity.Product{}).Where("product_id = ?", id).Update("product_description", request.D).Error; err != nil {
			ctx.JSON(400, gin.H{"message": "Error"})
			return
		}

		ctx.JSON(200, gin.H{"message": "Product Description updated"})
	}
}

func EditProductCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			Cat string `json:"product_category"`
		}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid request"})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Print(err)
			ctx.JSON(404, gin.H{"message": "No such product"})
			return
		}

		if err := db.Model(&entity.Product{}).Where("product_id = ?", id).Update("product_category", request.Cat).Error; err != nil {
			ctx.JSON(400, gin.H{"message": "Error"})
			return
		}

		ctx.JSON(200, gin.H{"message": "Product Category updated"})
	}
}

func EditProductQuantity() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			Quantity int `json:"product_quantity"`
		}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid request"})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Print(err)
			ctx.JSON(404, gin.H{"message": "No such product"})
			return
		}

		if err := db.Model(&entity.Product{}).Where("product_id = ?", id).Update("product_quantity", request.Quantity).Error; err != nil {
			ctx.JSON(400, gin.H{"message": "Error"})
			return
		}

		ctx.JSON(200, gin.H{"message": "Product Quantity updated"})
	}
}

func EditProductStockQuantity() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			Quantity int `json:"product_stock_quantity"`
		}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid request"})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Print(err)
			ctx.JSON(404, gin.H{"message": "No such product"})
			return
		}

		if err := db.Model(&entity.Product{}).Where("product_id = ?", id).Update("product_stock_quantity", request.Quantity).Error; err != nil {
			ctx.JSON(400, gin.H{"message": "Error"})
			return
		}

		ctx.JSON(200, gin.H{"message": "Product Stock Quantity updated"})
	}
}

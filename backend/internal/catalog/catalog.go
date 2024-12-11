package controllers

import (
	"database/sql"
	"fmt"
	config "myapp/internal/data_base"
	entity "myapp/internal/structures"
	"net/http"
	"strconv"
	"log"
//	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitCatalog(DB *gorm.DB, eng *gin.Engine) {
	db = DB

	eng.GET("/catalog", GetCatalogItems)

	eng.POST("/filter", ProductFilter)
	eng.GET("/fav_items", GetFavoriteItems)
	eng.GET("/get_item_page/:id", GetItem)
	eng.POST("/fav_items/:id", AddToFavorites)
	eng.DELETE("/fav_items/:id", RemoveFromFavorites)
	eng.POST("/cart/add/:id", AddToCart)
	eng.DELETE("/cart/remove/:id", RemoveFromCart)
	eng.GET("/cart", GetCart)
	eng.GET("/cart/:id", GetItem)

	eng.POST("/serch_item/name", SerchProductsByName())
}

func GetCatalogItems(c *gin.Context) {
	var products []entity.Product
	if err := db.Model(&entity.Product{}).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Недействительный ID продукта"})
		return
	}

	db, err := sql.Open("mysql", "admin_for_itam_store:your_password@tcp(147.45.163.58:3306)/itam_store")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Printf("Подключено")
	//Установка данных
	//insert, err := db.Query(fmt.Sprintf("INSERT INTO test.articles (`title`, `anons`, `full_text`) VALUES ('%s', '%s', '%s')", title, anons, full_text))
	var zapros = fmt.Sprintf("SELECT product_name, product_price, product_description, product_category FROM `products` WHERE product_id  = '%d'", id)
	fmt.Println(zapros)
	res, err := db.Query(zapros)
	var product entity.Product
	for res.Next() {

		err = res.Scan(&product.Name, &product.Price, &product.Description, &product.Category)

	}

	var zapros2 = fmt.Sprintf("SELECT id_feature, value FROM `added_features` WHERE id_item  = '%d'", id)
	fmt.Println(zapros2)
	res2, err := db.Query(zapros2)
	var features []entity.Feature
	for res2.Next() {
		var tempFeature entity.Feature
		err = res2.Scan(&tempFeature.Name, &tempFeature.Value)
		features = append(features, tempFeature)

	}
	combined := map[string]interface{}{
        "product":    product,
        "features": features,
    }
	//product.Features = features

	//	if err := config.DB.First(&product, id).Error; err != nil {
	//		fmt.Println("err")
	//		panic(err)
	///		c.JSON(http.StatusInternalServerError, gin.H{"Ошибка": "Продукт не найден3"})

	//	return
	//}
	//fmt.Println(product)
	c.JSON(http.StatusOK, combined)
}

func GetFavoriteItems(c *gin.Context) {
	//session := sessions.Default(c)
	userID := 26 //session.Get("id")
	//if userID == nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{"Ошибка": "Пользователь не авторизован"})
	//	return
	//}

	var favorites_ids []int
	query := "SELECT product_id FROM favourites WHERE user_id = ?"
	result := db.Raw(query, userID).Scan(&favorites_ids)
	if result.Error != nil {
		log.Print(result.Error)
		c.JSON(500, gin.H{"message": "failed get favourities"})
		return
	}

	var favorites []entity.Product
	query = "SELECT * FROM products WHERE product_id IN (?)"
	result = db.Raw(query, favorites_ids).Scan(&favorites)
	if result.Error != nil {
		log.Print(result.Error)
		c.JSON(500, gin.H{"message": "failed get favourities"})
		return
	}

	c.JSON(http.StatusOK, favorites)
}

func AddToFavorites(c *gin.Context) {
	//session := sessions.Default(c)
	userID := 26 //session.Get("id")
	//if userID == nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{"Ошибка": "Пользователь не авторизован"})
	//	return
	//}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Недействительный ID продукта"})
		return
	}

	query := "INSERT INTO favourites (user_id, product_id) VALUES (?, ?)"
	result := db.Exec(query, userID, id)
	if result.Error != nil {
		log.Print(result.Error)
		c.JSON(500, gin.H{"message": "Failed add to favourites"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Сообщение": "Продукт добавлен в избранное"})
}

func RemoveFromFavorites(c *gin.Context) {
	//session := sessions.Default(c)
	userID := 26 //session.Get("id")
	//if userID == nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{"Ошибка": "Пользователь не авторизован"})
	//	return
	//}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Недействительный ID продукта"})
		return

	}

	query := "DELETE FROM favourites WHERE user_id = ? AND product_id = ?"
	result := db.Exec(query, userID, id)
	if result.Error != nil {
		log.Print(result.Error)
		c.JSON(500, gin.H{"message": "Failed remove"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Сообщение": "Продукт удален из избранного"})
}

func ProductFilter(c *gin.Context) {
	var filterParams entity.FilterParams
	if err := c.ShouldBindJSON(&filterParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}

	var products []entity.Product
	query := config.DB

	if filterParams.Category != "" {
		query = query.Where("product_category = ?", filterParams.Category)
	}

	query = query.Where("product_price >= ? AND product_price <= ?", filterParams.MinPrice, filterParams.MaxPrice)

	if err := query.Find(&products).Error; err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetCart(c *gin.Context) {
	//session := sessions.Default(c)
	userID := 26 //session.Get("id")
	//if userID == nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{"Ошибка": "Пользователь не авторизован"})
	//	return
	//}

	var cart_ids []int
	query := "SELECT product_id FROM user_cart WHERE user_id = ?"
	result := db.Raw(query, userID).Scan(&cart_ids)
	if result.Error != nil {
		log.Print(result.Error)
		c.JSON(500, gin.H{"message": "failed get cart"})
		return
	}

	var cart []entity.Product
	query = "SELECT * FROM products WHERE product_id IN (?)"
	result = db.Raw(query, cart_ids).Scan(&cart)
	if result.Error != nil {
		log.Print(result.Error)
		c.JSON(500, gin.H{"message": "failed get cart"})
		return
	}

	c.JSON(http.StatusOK, cart)
}

func AddToCart(c *gin.Context) {
	//sessions := sessions.Default(c)
	userID := 26 //sessions.Get("id")
	//if id == nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Недействительный ID пользователя"})
	//	return
	//}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Недействительный ID продукта"})
		return
	}

	query := "INSERT INTO user_cart (user_id, product_id) VALUES (?, ?)"
	result := db.Exec(query, userID, id)
	if result.Error != nil {
		log.Print(result.Error)
		c.JSON(500, gin.H{"message": "Failed add to cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Сообщение": "Продукт добавлен в корзину"})
}

func RemoveFromCart(c *gin.Context) {
	//session := sessions.Default(c)
	userID := 26 //session.Get("id")
	//if userID == nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{"Ошибка": "Пользователь не авторизован"})
	//	return
	//}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Недействительный ID продукта"})
		return

	}

	query := "DELETE FROM user_cart WHERE user_id = ? AND product_id = ?"
	result := db.Exec(query, userID, id)
	if result.Error != nil {
		log.Print(result.Error)
		c.JSON(500, gin.H{"message": "Failed remove"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Сообщение": "Продукт удален из карзины"})
}

func SerchProductsByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type request struct {
			Name string `json:"product_name"`
		}
		var r request

		if err := ctx.ShouldBindJSON(&r); err != nil {
			ctx.JSON(400, gin.H{"message": "Bad request"})
			return
		}

		var products []entity.Product

		query := "SELECT * FROM products WHERE product_name LIKE ?"
		searchPattern := "%" + r.Name + "%"

		result := db.Raw(query, searchPattern).Scan(&products)
		if result.Error != nil {
			log.Print(result.Error)
			ctx.JSON(500, gin.H{"message": "Failed get products"})
			return
		}
		ctx.JSON(200, products)
	}
}

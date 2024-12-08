package controllers

import (
	"database/sql"
	"fmt"
	config "myapp/internal/data_base"
	entity "myapp/internal/structures"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
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
	eng.GET("/fav_items/:id", GetItem)
	eng.POST("/fav_items/:id", AddToFavorites)
	eng.DELETE("/fav_items/:id", RemoveFromFavorites)
	eng.POST("/cart/add/:id", AddToCart)
	eng.DELETE("/cart/remove/:id", RemoveFromCart)
	eng.GET("/cart", GetCart)
	eng.GET("/cart/:id", GetItem)
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
	//product.Features = features

	//	if err := config.DB.First(&product, id).Error; err != nil {
	//		fmt.Println("err")
	//		panic(err)
	///		c.JSON(http.StatusInternalServerError, gin.H{"Ошибка": "Продукт не найден3"})

	//	return
	//}
	fmt.Println(product)
	c.JSON(http.StatusOK, product)
}

func GetFavoriteItems(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Ошибка": "Пользователь не авторизован"})
		return
	}

	var favorites []entity.Favorite
	if err := config.DB.Where("user_id = ?", userID).Find(&favorites).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
		return
	}

	var favoriteProductIDs []uint
	for _, fav := range favorites {
		favoriteProductIDs = append(favoriteProductIDs, fav.ProductID)
	}

	var favoriteProducts []entity.Product
	if len(favoriteProductIDs) > 0 {
		if err := config.DB.Where("id IN ?", favoriteProductIDs).Find(&favoriteProducts).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, favoriteProducts)
}

func AddToFavorites(c *gin.Context) {

	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Ошибка": "Пользователь не авторизован"})
		return
	}

	var favorite entity.Favorite
	if err := c.ShouldBindJSON(&favorite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}
	if err := config.DB.Create(&favorite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Сообщение": "Продукт добавлен в избранное"})

}

func RemoveFromFavorites(c *gin.Context) {

	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Ошибка": "Пользователь не авторизован"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Недействительный ID продукта"})
		return

	}
	if err := config.DB.Where("id = ?", id).Delete(&entity.Favorite{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
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
		query = query.Where("category = ?", filterParams.Category)
	}

	query = query.Where("price >= ? AND price <= ?", filterParams.MinPrice, filterParams.MaxPrice)

	if err := query.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
		return

	}

	c.JSON(http.StatusOK, products)

}
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

	sessions := sessions.Default(c)
	id := sessions.Get("id")
	if id == nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Недействительный ID пользователя"})
		return
	}

	var cartItem entity.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})

		return
	}

	if err := config.DB.Create(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"Сообщение": "Продукт добавлен в корзину"})
}

func RemoveFromCart(c *gin.Context) {

	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Ошибка": "Пользователь не авторизован"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Недействительный ID продукта"})

		return
	}

	if err := config.DB.Where("id= ?", id).Delete(&entity.CartItem{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"Сообщение": "Продукт удален из карзины"})
}

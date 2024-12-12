package controllers

import (
	"database/sql"
	"fmt"
	"log"
	config "myapp/internal/data_base"
	entity "myapp/internal/structures"
	"net/http"
	"strconv"

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
	for i := 0; i < len(products); i++ {
		product := products[i]
		UserID := 26
		query := "SELECT EXISTS(SELECT 1 FROM  favourites WHERE  product_id = ? AND user_id = ?)"
		var exists int
		_ = db.Raw(query, product.ProductID, UserID).Scan(&exists)

		products[i].Is_in_fav = exists

		query = "SELECT EXISTS(SELECT 1 FROM  user_cart WHERE  product_id = ? AND user_id = ?)"

		_ = db.Raw(query, product.ProductID, UserID).Scan(&exists)

		products[i].Is_in_cart = exists
	}

	c.JSON(http.StatusOK, products)
}

func GetItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Недействительный ID продукта"})
		return
	}

	var product entity.Product
	UserID := 26
	query := "SELECT EXISTS(SELECT 1 FROM  favourites WHERE  product_id = ? AND user_id = ?)"
	var exists int
	_ = db.Raw(query, id, UserID).Scan(&exists)

	product.Is_in_fav = exists

	query = "SELECT EXISTS(SELECT 1 FROM  user_cart WHERE  product_id = ? AND user_id = ?)"

	_ = db.Raw(query, id, UserID).Scan(&exists)

	product.Is_in_cart = exists

	db, err := sql.Open("mysql", "admin_for_itam_store:your_password@tcp(147.45.163.58:3306)/itam_store")
	if err != nil { // Подклчение к бд для работы с Query
		panic(err)
	}

	defer db.Close()
	fmt.Printf("Подключено")

	var zapros = fmt.Sprintf("SELECT product_name, product_price, product_description, product_category, product_image FROM `products` WHERE product_id  = '%d'", id)
	fmt.Println(zapros) // Получение всей информации о продукте по его id
	res, err := db.Query(zapros)

	for res.Next() {

		err = res.Scan(&product.Name, &product.Price, &product.Description, &product.Category, &product.Image)

	}

	var zapros2 = fmt.Sprintf("SELECT id_feature, value FROM `added_features` WHERE id_item  = '%d'", id)
	fmt.Println(zapros2) // Получение всех id фич и их значений по id продукта
	res2, err := db.Query(zapros2)
	var features []entity.Feature
	for res2.Next() {
		var tempFeature entity.Feature
		var id_fetur string
		err = res2.Scan(&id_fetur, &tempFeature.Value)

		var zapros3 = fmt.Sprintf("SELECT name, unit_of_measurement FROM `specifications_for_item` WHERE 	id   = '%s'", id_fetur)
		fmt.Println(zapros3) // Получение всех id фич и их значений по id продукта
		res3, _ := db.Query(zapros3)
		fmt.Println(tempFeature)
		for res3.Next() {

			err = res3.Scan(&tempFeature.Name, &tempFeature.Unit_of_measurement)
		}
		fmt.Println(tempFeature)
		features = append(features, tempFeature)

	}
	combined := map[string]interface{}{
		"product":  product,
		"features": features,
	} // Формирование json файла. Сам продукт и массив фич к нему

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

	query := "SELECT EXISTS(SELECT 1 FROM  favourites WHERE user_id = ? AND product_id = ?)"
	var exists bool
	result := db.Raw(query, userID, id).Scan(&exists)

	if result.Error != nil {
		// Обработка ошибки
		fmt.Println("Ошибка при выполнении запроса:", result.Error)
	} else {
		fmt.Println("Существует ли продукт:", exists)
	}

	if exists {
		c.JSON(http.StatusOK, gin.H{"Сообщение": "Был ранее добавлен"})
		return
	}

	query = "INSERT INTO favourites (user_id, product_id) VALUES (?, ?)"
	result = db.Exec(query, userID, id)
	if result.Error != nil {
		log.Print(result.Error)
		c.JSON(500, gin.H{"message": "Failed add to favourites"})
		return
	}

	query = "UPDATE products SET is_in_fav = ? WHERE product_id = ?"
	result = db.Exec(query, true, id)
	if result.Error != nil {
		c.JSON(500, gin.H{"message": "cant add"})
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

	query = "UPDATE products SET is_in_fav = ? WHERE product_id = ?"
	result = db.Exec(query, false, id)
	if result.Error != nil {
		c.JSON(500, gin.H{"message": "cant add"})
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

	query := "SELECT EXISTS(SELECT 1 FROM user_cart WHERE user_id = ? AND product_id = ?)"
	var exists bool
	result := db.Raw(query, userID, id).Scan(&exists)

	if result.Error != nil {
		// Обработка ошибки
		fmt.Println("Ошибка при выполнении запроса:", result.Error)
	} else {
		fmt.Println("Существует ли продукт:", exists)
	}

	query = "UPDATE products SET is_in_cart = ? WHERE product_id = ?"
	result = db.Exec(query, true, id)
	if result.Error != nil {
		c.JSON(500, gin.H{"message": "cant add"})
		return
	}

	fmt.Println(exists)
	if exists {
		c.JSON(http.StatusOK, gin.H{"Сообщение": "Был ранее добавлен"})
		return
	}

	query = "INSERT INTO user_cart (user_id, product_id) VALUES (?, ?)"
	result = db.Exec(query, userID, id)
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

	query = "UPDATE products SET is_in_cart = ? WHERE product_id = ?"
	result = db.Exec(query, false, id)
	if result.Error != nil {
		c.JSON(500, gin.H{"message": "cant add"})
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

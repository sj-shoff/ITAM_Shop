# ITAM_Shop


### Эндпоинты
Сценарии:<br />
  Добавление продукта в магазин<br />
    -1. /add_item На вход подается JSON файл по подобию струкутры entity.Product. Возвращается ответ в виде json строчки с результатом добавления (Успешно или нет) <br />
  Добавление\удаление товара в корзину (необходима авторизация)<br />
    -1. POST /cart/item/:id Передается Id товара для добавления в коризну. Возвращается JSON строка с результатом (Успешно или нет) <br />
    -2. DELETE /cart/item/:id Передается Id товара для удаление из коризны. Возвращается JSON строка с результатом (Успешно или нет) <br />
    
### Используемые сущности 


```



# Product
```bash
type Product struct {
	gorm.Model
	ProductID      uint    `json:"product_id"`
	Price          float64 `json:"product_price"`
	Name           string  `json:"product_name"`
	ImageURL       string  `json:"image_url"`
	Description    string  `json:"product_description"`
	Category       string  `json:"product_category"`
	Specifications string  `json:"product_specifications"`
	Quantity       int     `json:"product_quantity"`
	StockQuantity  int     `json:"product_stock_quantity"`
	//
}
```


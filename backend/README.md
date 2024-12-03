# ITAM_Shop Backend

ITAM_Shop - это веб-приложение, разработанное с использованием Go и фреймворка Gin. Эта часть репозитория содержит серверную часть приложения, которая обрабатывает запросы, управляет сессиями и взаимодействует с базой данных.

### Предварительные требования

- Установите [Go]
- Убедитесь что с вашего компьютера есть доступ к удаленной базе данных. Если такогого нет обращаться к @BRDDRTy в Telegram

### Клонирование репозитория

```bash
git clone https://github.com/sj-shoff/ITAM_Shop.git
```
### Запуск проекта

- Переключитесь на нужную ветку(подразумевается develop)
```bash
git checkout develop
```

- Перейдите в директорию backend
```bash
cd backend/
```

- Проверьте установленные зависимости
```bash
go mod tidy
```

- Для запуска пропишите следущие команды
```bash
cd cmd
go run .
```
Сервер будет запущен на порту http://localhost:8080
Проверить cтатус сервера можно по http://localhost:8080/health

### Эндпоинты
Здоровье сервера

    GET /health - Проверка состояния сервера.

Регистрация и вход

    POST /register - Добавляет пользователя в базу данных и сохраняет в cookie код подтверждения почты Request -> entity.User Response -> .
    POST /login    - Вход пользователя в аккаунт.
    POST /checkemail - Отправка email с кодом и его подтверждение Request -> {"code":"..."}. Response -> 200 если код правильный
    POST /newpassword - Обновление пароля если пользователь зарегистрирован или восстановление пароль Request -> `json:"user_password"`
    POST /recoverpassword - Восстановление пароля Request -> `json:"user_login"`
    
    Сценарии:
        Регистрация:
        - 1. /register Отправка запроса на регистрацию формата entity.User(user_login, user_email, user_password) 
        - 2. Со стороны бэка отправляется письмо пользователю с кодом подтверждения после чего пользователь должен подтвердить почту
        - 3. /checkemail Запрос -> `json:"code"`. Подтверждение почты Response 200 OK
        - 4. После получения ответа пользователь успешно зарегистрирован
        Вход:
        - 1. /login Вход пользователя 
        Восстановление пароля:
        - 1. /recoverpassword Отправка запроса `json:"user_login"`
        - 2. Со стороны бэка отправляется письмо пользователю с кодом подтверждения после чего пользователь должен подтвердить почту
        - 3. /checkemail Запрос -> `json:"code"`. Подтверждение почты Response 200 OK
        - 4. После подтверждения почты обновление пароля
        - 5. /newpassword Запрос -> `json:"user_password"`
        - 6. После получения ответа 200 пароль успешно изменен

Личный кабинет

    POST /logout - выход из личного кабинета
    POST /updateavatar - Обновить аватар пользователя
    POST /updatename - Обновить имя пользователя
    POST /updatesurname - Обновить фамилию пользователя
    POST /updatepassword - Обновить пароль пользователя


Транзакции

    POST /givemoney Request {`json:"user_login"`, `json:"user_money"`} Добавить деньги на счет пользователя

    Для бэкэнда:
        - Добавлена функция TakeOffMoney(login, price) -> bool Возвращает прошла ли оплата или нет

Каталог 

    
	GET /catalog - Получение списка товаров из каталога
	POST /filter - Применение фильтра к товарам в каталоге

	GET /fav_items - Получение списка товаров, добавленных в избранное
	GET /fav_items/:id - Получение страницы товара по его id
	POST /fav_items/:id - Добавление товара в избранное 
	DELETE /fav_items/:id - Удаление товара из избранного
	
	GET /cart - Получение списка товаров в корзине
    GET /cart/:id - Получение страницы товара в корзине по его id 
	POST /cart/add/:id - Добавление товара в корзину
	DELETE /cart/remove/:id - Удаление товара из корзины

	


    

Административные функции

    POST /createnewproduct - создание нового продукта(без логики администратора) Request -> entity.Product
    POST /editproduct/:id - редактирование продукта(без логики администратора) Request -> entity.Product
    POST /deleteproduct/:id - удаление продукта(без логики администратора)


    NOT OK

    GET /analytics - Получение аналитики.
    GET /admin_panel - Доступ к административной панели.
    POST /giveadminrights - Выдача прав доступа админа по логину Request -> {`json:"user_login"`}


### Используемые сущности 

# User
```bash
type User struct {
	ID          uint    `json:"user_id" gorm:"column:user_id"`
	Balance     float64 `json:"user_balance" gorm:"column:user_balance"`
	UserName    string  `json:"user_name" gorm:"column:user_name"`
	UserSurname string  `json:"user_surname" gorm:"column:user_surname"`
	Login       string  `json:"user_login" gorm:"column:user_login"`
	Email       string  `json:"user_email" gorm:"column:user_email"`
	Password    string  `json:"user_password" gorm:"column:user_password"`
	Admin       bool    `json:"user_admin_rights" gorm:"column:user_admin_rights"`
	Avatar      []byte  `json:"user_avatar" gorm:"column:user_avatar"`
	//
}
```

# FilterParams
```bash
type FilterParams struct {
	Category string  `json:"category"`
	MinPrice float64 `json:"minPrice"`
	MaxPrice float64 `json:"maxPrice"`
}
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

# CartItem
```bash
type CartItem struct {
	ProductID_cart uint    `json:"product_id"`
	Quantity       int     `json:"quantity"`
	Product        Product `json:"product"`
}
```

# Cart
```bash
type Cart struct {
	Items []CartItem `json:"items"`
}
```

# Favourite
```bash
type Favorite struct {
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
	IsFavorite     bool    `json:"is_favorite"`
}
```

# Order
```bash
type Order struct {
	gorm.Model
	UserID uint   `gorm:"not null" json:"user_id"`
	Status string `gorm:"not null" json:"status"`
}
```

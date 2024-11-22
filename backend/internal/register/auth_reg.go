package main

//Файл отвечающий за авторизация и регистрацию пользователей
// func CreateUser(user *entity.User) error {
// 	stmt, err := db.Prepare("INSERT INTO users(username, password) VALUES(?, ?)")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()
// 	fmt.Println("oki")
// 	_, err = stmt.Exec(user.Username, user.Password)
// 	fmt.Println("NOToki")
// 	return err
// }

// func ShowRegistrationForm(c *gin.Context) {
// 	c.HTML(http.StatusOK, "register.html", nil)
// }

// func RegisterUser(c *gin.Context) {
// 	var user entity.User

// 	if err := c.ShouldBind(&user); err != nil {
// 		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "Invalid form data"})
// 		return
// 	}

// 	// Валидация данных
// 	if user.Username == "" || user.Password == "" {
// 		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "Username and password are required"})
// 		return
// 	}

// 	// Хеширование пароля
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		c.HTML(http.StatusInternalServerError, "register.html", gin.H{"error": "Failed to hash password"})
// 		return
// 	}
// 	user.Password = string(hashedPassword)

// 	// Сохранение пользователя в базе данных
// 	if err := CreateUser(&user); err != nil {
// 		fmt.Println(err)
// 		c.HTML(http.StatusInternalServerError, "register.html", gin.H{"error": "Failed to create user"})
// 		return
// 	}

// 	c.HTML(http.StatusOK, "register.html", gin.H{"success": "Registration successful"})
// }

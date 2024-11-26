package main

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	entity "myapp/internal/structures"
	"net/http"
	"net/smtp"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db    *sql.DB
	store cookie.Store
	s     *gin.Engine
)

func userExists(login string) bool {
	var exists bool = false
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE user_login = ?)"

	err := db.QueryRow(query, login).Scan(&exists)
	if err != nil {
		log.Print(err)
		return false
	}
	return exists
}

func generateCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(1000000)
	return fmt.Sprintf("%06d", code)
}

func ConfirmEmail(email string) string {
	// Данные SMTP сервера
	smtpHost := "smtp.gmail.com"
	smtpPort := "465"
	smtpUser := "itamshophelp@gmail.com"
	smtpPass := "jdazutaivdfequqh"

	// Формирование письма
	code := generateCode()
	subject := "Подтверждение электронной почты"
	body := fmt.Sprintf("Пожалуйста, подтвердите вашу почту. Код для подтверждения: %s", code)
	msg := []byte("To: " + email + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	// Установка соединения с SMTP сервером
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}
	conn, err := tls.Dial("tcp", smtpHost+":"+smtpPort, tlsConfig)
	if err != nil {
		fmt.Println("Ошибка при подключении к SMTP серверу:", err)
		return ""
	}
	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		fmt.Println("Ошибка при создании клиента SMTP:", err)
		return ""
	}

	// Аутентификация
	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
	if err = c.Auth(auth); err != nil {
		fmt.Println("Ошибка аутентификации:", err)
		return ""
	}

	// Отправка письма
	if err = c.Mail(smtpUser); err != nil {
		fmt.Println("Ошибка при указании отправителя:", err)
		return ""
	}
	if err = c.Rcpt(email); err != nil {
		fmt.Println("Ошибка при указании получателя:", err)
		return ""
	}
	w, err := c.Data()
	if err != nil {
		fmt.Println("Ошибка при получении объекта для записи данных:", err)
		return ""
	}
	_, err = w.Write(msg)
	if err != nil {
		fmt.Println("Ошибка при записи данных:", err)
		return ""
	}
	err = w.Close()
	if err != nil {
		fmt.Println("Ошибка при закрытии объекта:", err)
		return ""
	}
	c.Quit()

	fmt.Println("Письмо успешно отправлено!")
	return code
}

func RegisterUSER() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user entity.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			log.Print("Failed bind JSON ", err)
			return
		}

		if userExists(user.Login) {
			ctx.JSON(204, gin.H{"message": "User with this login is already exists"})
			return
		}

		_, err := db.Exec("INSERT into users (user_login, user_email, user_password) VALUES (?, ?, ?)", user.Login, user.Email, user.Password)
		if err != nil {
			log.Print("Failed insert user ", err)
			ctx.JSON(400, gin.H{"message": "Error"})
			return
		}

		code := ConfirmEmail(user.Email)

		sessions := sessions.Default(ctx)
		sessions.Set("login", user.Login)
		sessions.Set("code", code)

		if err := sessions.Save(); err != nil {
			log.Print("Failed to save session: ", err)
			ctx.JSON(400, gin.H{"message": "Error"})
			return
		}

		ctx.JSON(200, gin.H{"message": "Please check your email"})
	}
}

func Checkemail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			Code string `json:"code"`
		}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid request"})
			return
		}

		sessions := sessions.Default(ctx)
		login := sessions.Get("login")
		code := sessions.Get("code")

		var v bool = true

		if code == request.Code {
			query := "UPDATE users SET user_verified_email = ? WHERE user_login = ?"
			_, err := db.Exec(query, v, login)
			if err != nil {
				log.Print("Failed to update verified_email: ", err)
				ctx.JSON(500, gin.H{"message": "Internal server error"})
				return
			}

			ctx.JSON(200, gin.H{"message": "Successful verified email"})

		} else {
			log.Print("wrong code")
			ctx.JSON(204, gin.H{"message": "wrong code"})
			return
		}
	}
}

func LoginUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user entity.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			log.Print("Failed bind json ", err)
			return
		}

		var userOK entity.User
		query := "SELECT user_password, user_email FROM users WHERE user_login = ?"

		err := db.QueryRow(query, user.Login).Scan(&userOK.Password, &userOK.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Print("No such user", err)
				return
			}
			log.Print("Failed connect to DB ", err)
			return
		}

		if user.Password != userOK.Password {
			ctx.JSON(204, gin.H{"message": "Wrong login or password"})
			return
		}

		sessions := sessions.Default(ctx)
		sessions.Set("login", user.Login)

		if err := sessions.Save(); err != nil {
			log.Print("Failed to save session: ", err)
		}

		ctx.JSON(200, gin.H{"message": "user logged", "mail": userOK.Email})
	}
}

func RecoverUserPassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			Login string `json:"user_login"`
		}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid request"})
			return
		}

		var userOK entity.User
		query := "SELECT user_password, user_email FROM users WHERE user_login = ?"

		err := db.QueryRow(query, request.Login).Scan(&userOK.Password, &userOK.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(400, gin.H{"message": "no such user"})
				return
			}
			log.Print("Failed connect to DB ", err)
			return
		}

		var v bool = false

		query = "UPDATE users SET user_verified_email = ? WHERE user_login = ?"
		db.Exec(query, v, request.Login)

		code := ConfirmEmail(userOK.Email)

		sessions := sessions.Default(ctx)
		sessions.Set("login", userOK.Login)
		sessions.Set("code", code)

		if err := sessions.Save(); err != nil {
			log.Print("Failed to save session: ", err)
		}

		ctx.JSON(200, gin.H{"message": ""})
	}
}

func Newpassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			New_password string `json:"user_password"`
		}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid request"})
			return
		}

		sessions := sessions.Default(ctx)
		login := sessions.Get("login")

		query := "UPDATE users SET user_password = ? WHERE user_login = ?"
		_, err := db.Exec(query, request.New_password, login)
		if err != nil {
			ctx.JSON(400, gin.H{"message": "Error"})
			log.Print("BD error")
			return
		}

		ctx.JSON(200, gin.H{"message": "Password updated"})
	}
}

func main() {
	fmt.Println("Starting server!")
	var err error

	db, err = sql.Open("mysql", "admin_for_itam_store:your_password@tcp(147.45.163.58:3306)/itam_store")
	if err != nil {
		log.Fatal("Failed open DB ", err)
	}

	defer db.Close()

	s = gin.Default()

	store = cookie.NewStore([]byte("secret-key"))
	s.Use(sessions.Sessions("mysession", store))

	s.GET("/health", func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusOK)
	})
	s.POST("/register", RegisterUSER())
	s.POST("/login", LoginUser())
	s.POST("/checkemail", Checkemail())
	s.POST("/recoverpassword", RecoverUserPassword())
	s.POST("/newpassword", Newpassword())

	s.Run(":8090")

	fmt.Println("Server is running on :8090")
}

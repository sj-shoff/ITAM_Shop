package main

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Login    string `json:"login"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type verifyUser struct {
	login string
	code  string
}

func userExists(db *sql.DB, username string) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)"

	err := db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func generateCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(1000000)
	return fmt.Sprintf("%06d", code)
}

func confirmEmail(email string) string {
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

func registerUSER(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			log.Print("Failed bind JSON ", err)
			return
		}

		if userExists(db, user.Login) {
			ctx.JSON(204, gin.H{"message": "User with this login is already exists"})
			return
		}

		_, err := db.Exec("INSERT into users (login, email, password) VALUES (?, ?, ?)", user.Login, user.Mail, user.Password)
		if err != nil {
			log.Print("Failed insert user ", err)
			return
		}

		code := confirmEmail(user.Mail)

		sessions := sessions.Default(ctx)
		sessions.Set("login", user.Login)
		sessions.Set("code", code)

		if err := sessions.Save(); err != nil {
			log.Print("Failed to save session: ", err)
		}

		login := sessions.Get("login")
		code1 := sessions.Get("code")

		log.Print(login, code1, code)

		ctx.JSON(200, gin.H{"message": "Please check your email"})
	}
}

func checkemail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			Code string `json:"message"`
		}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid request"})
			return
		}

		sessions := sessions.Default(ctx)
		login := sessions.Get("login")
		code := sessions.Get("code")

		log.Printf("Login: %s, Code: %s, Provided Code: %s", login, code, request.Code)

		if code == request.Code {
			ctx.JSON(200, gin.H{"message": "Successful verified email"})
			query := "UPDATE users SET verified_email = ? WHERE login = ?"
			_, err := db.Exec(query, true, login)
			if err != nil {
				log.Print("Failed to update verified_email: ", err)
				ctx.JSON(500, gin.H{"message": "Internal server error"})
				return
			}
		} else {
			log.Printf("fghjkl")
			ctx.JSON(204, gin.H{"message": "wrong code"})
			return
		}
	}
}

func loginUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			log.Print("Failed bind json ", err)
		}

		var userOK User
		query := "SELECT password, mail FROM users WHERE login = ?"

		// Выполняем запрос
		err := db.QueryRow(query, user.Login).Scan(&userOK.Password, &userOK.Mail)
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

		ctx.JSON(200, gin.H{"message": "user logged", "mail": userOK.Mail})
	}
}

func recoverUserPassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		login, ok := ctx.Get("login")
		if !ok {
			ctx.JSON(400, gin.H{"message": "wrong login"})
		}

		var userOK User
		query := "SELECT password, mail FROM users WHERE login = ?"

		err := db.QueryRow(query, login).Scan(&userOK.Password, &userOK.Mail)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(400, gin.H{"message": "no such user"})
				return
			}
			log.Print("Failed connect to DB ", err)
			return
		}

		query = "UPDATE users SET verified_email = ? WHERE login = ?"
		db.Exec(query, login, false)

		code := confirmEmail(userOK.Mail)

		sessions := sessions.Default(ctx)
		sessions.Set("login", userOK.Login)
		sessions.Set("code", code)

		if err := sessions.Save(); err != nil {
			log.Print("Failed to save session: ", err)
		}

		ctx.JSON(200, gin.H{"message": ""})
	}
}

func newpassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessions := sessions.Default(ctx)
		login := sessions.Get("login")

		query := "UPDATE users SET password = ? WHERE login = ?"
		db.Exec(query, login, true)

	}
}

var (
	db    *sql.DB
	store cookie.Store
	s     *gin.Engine
)

func main() {
	fmt.Println("Starting server!")

	db, err := sql.Open("mysql", "admin_for_itam_store:your_password@tcp(147.45.163.58:3306)/itam_store")
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
	s.POST("/register", registerUSER(db))
	s.POST("/login", loginUser(db))
	s.POST("/checkemail", checkemail())
	s.POST("/recoverpassword", recoverUserPassword())
	//s.POST("/newpassword", newpassword())

	s.Run(":8090")

	fmt.Println("Server is running on :8090")

}

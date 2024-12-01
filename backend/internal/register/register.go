package register

import (
	"crypto/tls"
	"fmt"
	"log"
	"math/rand"
	entity "myapp/internal/structures"
	"net/smtp"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/gin-contrib/sessions"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *gorm.DB
)

func InitRegister(db1 *gorm.DB, s *gin.Engine) {

	db = db1

	s.POST("/register", RegisterUSER())
	s.POST("/login", LoginUser())
	s.POST("/checkemail", Checkemail())
	s.POST("/recoverpassword", RecoverUserPassword())
	s.POST("/newpassword", Newpassword())
	s.POST("/giveadminrights", GiveAdminRights())
}

func userExists(login string) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE user_login = ?)"
	err := db.Raw(query, login).Scan(&exists).Error
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

		query := "INSERT INTO users (user_login, user_email, user_password) VALUES (?, ?, ?)"
		if err := db.Exec(query, user.Login, user.Email, user.Password).Error; err != nil {
			log.Print("Не удалось вставить пользователя: ", err)
			ctx.JSON(400, gin.H{"message": "Ошибка"})
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
			if err := db.Model(&entity.User{}).Where("user_login = ?", login).Update("user_verified_email", v).Error; err != nil {
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

		if err := db.Raw(query, user.Login).Scan(&userOK).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				log.Print("Пользователь не найден: ", err)
				ctx.JSON(204, gin.H{"message": "No such user"})
				return
			}
			log.Print("Не удалось подключиться к базе данных: ", err)
			ctx.JSON(500, gin.H{"message": "Cant connect to DB"})
			return
		}

		if user.Password != userOK.Password {
			//log.Print(user.Password, userOK.Password)
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
		if err := db.Model(&userOK).Select("user_password, user_email").Where("user_login = ?", request.Login).First(&userOK).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(400, gin.H{"message": "no such user"})
				return
			}
			log.Print("Failed to connect to DB: ", err)
			return
		}

		var v bool = false

		if err := db.Model(&entity.User{}).Where("user_login = ?", request.Login).Update("user_verified_email", v).Error; err != nil {
			log.Print("Failed to update user_verified_email: ", err)
			ctx.JSON(500, gin.H{"message": "Internal server error"})
			return
		}

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
			NewPassword string `json:"user_password"`
		}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid request"})
			return
		}

		sessions := sessions.Default(ctx)
		login := sessions.Get("login")

		if err := db.Model(&entity.User{}).Where("user_login = ?", login).Update("user_password", request.NewPassword).Error; err != nil {
			ctx.JSON(400, gin.H{"message": "Error"})
			log.Print("Database error: ", err)
			return
		}

		ctx.JSON(200, gin.H{"message": "Password updated"})
	}
}

func GiveAdminRights() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			Login string `json:"user_login"`
		}

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid request"})
			return
		}

		if err := db.Model(&entity.User{}).Where("user_login = ?", request.Login).Update("user_admin_rights", true).Error; err != nil {
			ctx.JSON(400, gin.H{"message": "Error"})
			log.Print("Database error: ", err)
			return
		}

		ctx.JSON(200, gin.H{"message": "Admin rights successfully given"})

	}
}

package register

import (
	"log"
	emailalerts "myapp/internal/email_alerts"
	entity "myapp/internal/structures"

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
		code := emailalerts.ConfirmEmail(user.Email)

		sessions := sessions.Default(ctx)
		sessions.Set("login", user.Login)
		sessions.Set("code", code)
		sessions.Set("id", user.ID)

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
		sessions.Set("id", user.ID)

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

		code := emailalerts.ConfirmEmail(userOK.Email)

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

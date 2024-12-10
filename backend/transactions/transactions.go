package transactions

import (
	"log"
	entity "myapp/internal/structures"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *gorm.DB
)

func InitTransaction(db1 *gorm.DB, s *gin.Engine) {
	db = db1

	s.POST("/givemoney", GimmeSomeMoney())
}

func GimmeSomeMoney() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			Login string `json:"user_login"`
			Money string `json:"user_money"`
		}

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid request"})
			return
		}

		moneyAmount, err := strconv.ParseFloat(request.Money, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid money amount"})
			return
		}

		var user entity.User
		if err := db.Where("user_login = ?", request.Login).First(&user).Error; err != nil {
			ctx.JSON(404, gin.H{"message": "User not found"})
			return
		}

		user.Balance += moneyAmount
		if err := db.Save(&user).Error; err != nil {
			ctx.JSON(500, gin.H{"message": "Error updating balance"})
			log.Print("Database error: ", err)
			return
		}

		ctx.JSON(200, gin.H{"message": "Money added successfully", "new_balance": user.Balance})

	}
}

func TakeOffMoney(login string, price float64) bool {
	var user entity.User
	if err := db.Where("user_login = ?", login).First(&user).Error; err != nil {
		return false
	}

	user.Balance -= price
	if user.Balance < 0 {
		return false
	}
	if err := db.Save(&user).Error; err != nil {
		log.Print("Database error: ", err)
		return false
	}

	return true
}

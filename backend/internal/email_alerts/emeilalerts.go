package emailalerts

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	config "myapp/internal/data_base"
	entity "myapp/internal/structures"
	"net/smtp"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.DB
)

const (
	// Данные SMTP сервера
	smtpHost = "smtp.gmail.com"
	smtpPort = "465"
	smtpUser = "itamshophelp@gmail.com"
	smtpPass = "jdazutaivdfequqh"
)

func generateCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(1000000)
	return fmt.Sprintf("%06d", code)
}

func sendEmail(email string, data *[]byte) bool {
	// Установка соединения с SMTP сервером
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}
	conn, err := tls.Dial("tcp", smtpHost+":"+smtpPort, tlsConfig)
	if err != nil {
		fmt.Println("Ошибка при подключении к SMTP серверу:", err)
		return false
	}

	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		fmt.Println("Ошибка при создании клиента SMTP:", err)
		return false
	}

	// Аутентификация
	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
	if err = c.Auth(auth); err != nil {
		fmt.Println("Ошибка аутентификации:", err)
		return false
	}

	// Отправка письма
	if err = c.Mail(smtpUser); err != nil {
		fmt.Println("Ошибка при указании отправителя:", err)
		return false
	}

	if err = c.Rcpt(email); err != nil {
		fmt.Println("Ошибка при указании получателя:", err)
		return false
	}
	w, err := c.Data()
	if err != nil {
		fmt.Println("Ошибка при получении объекта для записи данных:", err)
		return false
	}

	_, err = w.Write(*data)
	if err != nil {
		fmt.Println("Ошибка при записи данных:", err)
		return false
	}
	err = w.Close()

	if err != nil {
		fmt.Println("Ошибка при закрытии объекта:", err)
		return false
	}
	c.Quit()
	fmt.Println("Письмо успешно отправлено!")
	return true
}

func ConfirmEmail(email string) string {
	// Формирование письма
	code := generateCode()
	subject := "Подтверждение электронной почты"
	body := fmt.Sprintf("Пожалуйста, подтвердите вашу почту. Код для подтверждения: %s", code)
	msg := []byte("To: " + email + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	if !sendEmail(email, &msg) {
		return ""
	}

	return code
}

func OrderReceipt(email string, ProductIDs []int) bool {
	// Отправка чека на почту

	var items []entity.Product

	if err := db.Where("id IN ?", ProductIDs).Find(&items).Error; err != nil {
		return false
	}

	var receipt strings.Builder
	receipt.WriteString("Чек:\n\n")
	var total float64

	for _, item := range items {
		receipt.WriteString(fmt.Sprintf("ID: %d\nНазвание: %s\nОписание: %s\nЦена: %.2f\n\n", item.ProductID, item.Name, item.Description, item.Price))
		total += item.Price
	}

	receipt.WriteString(fmt.Sprintf("Итого: %.2f\n", total))

	var msg []byte
	str := receipt.String()
	msg = []byte(str)

	return sendEmail(email, &msg)
}

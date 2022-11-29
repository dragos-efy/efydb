package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"strings"

	"github.com/efydb/config"
	"github.com/efydb/entities"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func isBlank(text string) bool {
	return len(strings.TrimSpace(text)) == 0
}

func ValidateUser(c *fiber.Ctx) (entities.User, error) {
	token := c.GetReqHeaders()["Authorization"]
	var user entities.User
	query := config.Database.Find(&user, &entities.User{
		Token: token,
	})
	if query.Error != nil {
		c.Status(fiber.StatusBadRequest).JSON(
			entities.Message{
				Message: "Invalid Access Token!",
			},
		)
	}

	return user, query.Error
}

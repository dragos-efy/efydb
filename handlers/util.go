package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/efydb/config"
	"github.com/efydb/entities"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
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

func ErrorResponse(c *fiber.Ctx, status int, error string) error {
	return c.Status(status).JSON(
		entities.Message{
			Message: error,
		},
	)
}

func OkResponse(c *fiber.Ctx) error {
	return c.Status(200).JSON(
		entities.Message{
			Message: "ok",
		},
	)
}

func SaveFile(c *fiber.Ctx, file *multipart.FileHeader) (string, error) {
	// create a unique name for the screenshot
	uniqueId := uuid.New()

	// generate a unique file name and extract the file extension
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]

	fullName := fmt.Sprintf("%s.%s", filename, fileExt)

	err := c.SaveFile(file, fmt.Sprintf("./files/%s", fullName))

	if err != nil {
		return "", err
	}

	url := c.BaseURL() + fmt.Sprintf("/files/%s", fullName)

	return url, nil
}

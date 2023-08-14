package util

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"strconv"
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

func IsBlank(text string) bool {
	return len(strings.TrimSpace(text)) == 0
}

func ValidateUser(c *fiber.Ctx) (entities.User, error) {
	var user entities.User
	token := c.Get("Authorization")

	if IsBlank(token) {
		ErrorResponse(c, fiber.StatusBadRequest, "Token can't be empty!")
		return user, errors.New("Token can't be empty!")
	}

	query := config.Database.Find(&user, &entities.User{
		Token: token,
	})
	if query.Error != nil {
		ErrorResponse(c, fiber.StatusBadRequest, "Invalid Access Token!")
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

	dir := fmt.Sprintf("%s/files/", config.RootDir())
	config.MkDirIfNotExists(dir)
	err := c.SaveFile(file, fmt.Sprintf("%s/%s", dir, fullName))

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("/files/%s", fullName), nil
}

func DeleteFile(path string) error {
	if IsBlank(path) {
		// the file doesn't need to be deleted if it doesn't exist
		return nil
	}

	fullPath := fmt.Sprintf("%s%s", config.RootDir(), path)
	return os.Remove(fullPath)
}

func ParseUintQuery(c *fiber.Ctx, key string) (uint, error) {
	query := c.Query(key)
	if query == "" {
		return 0, errors.New(fmt.Sprintf("No %s specified", key))
	}
	id, err := strconv.ParseUint(query, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func ParseUintParam(c *fiber.Ctx, key string) (uint, error) {
	param := c.Params(key)
	if param == "" {
		return 0, errors.New(fmt.Sprintf("No %s specified", key))
	}
	id, err := strconv.ParseUint(param, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

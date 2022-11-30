package handlers

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/efydb/config"
	"github.com/efydb/entities"
	"github.com/gofiber/fiber/v2"
)

func GetThemes(c *fiber.Ctx) error {
	showUnapproved := c.Query("showUnapproved", "false")

	var themes []entities.Theme

	if showUnapproved == "true" {
		config.Database.Find(&themes)
	} else {
		config.Database.Where("approved = ?", true).Find(&themes)
	}

	return c.JSON(&themes)
}

func CreateTheme(c *fiber.Ctx) error {
	// get the user
	user, err := ValidateUser(c)
	if err != nil {
		return nil
	}

	// get the uploaded screenshot
	ss, err := c.FormFile("screenshot")

	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "Screenshot missing!")
	}

	conf, err := c.FormFile("config")

	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "No config provided!")
	}

	data := c.FormValue("data", "")

	if data == "" {
		return ErrorResponse(c, fiber.StatusBadRequest, "Data can't be empty!")
	}

	var theme entities.Theme
	reader := strings.NewReader(data)
	jsonErr := json.NewDecoder(reader).Decode(&theme)

	if jsonErr != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	// create the new files and save them
	theme.Config, err = SaveFile(c, conf)
	if err != nil {
		return ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	theme.Screenshot, err = SaveFile(c, ss)
	if err != nil {
		return ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	// set approved false by default
	theme.Approved = false
	theme.Username = user.Name
	theme.Uploaded = time.Now().Unix()

	config.Database.Create(&theme)

	return c.Status(fiber.StatusCreated).JSON(theme)
}

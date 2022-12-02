package handlers

import (
	"net/url"
	"strings"
	"time"

	"github.com/efydb/config"
	"github.com/efydb/entities"

	"github.com/goccy/go-json"
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

	for index := range themes {
		rewriteTheme(&themes[index], c.BaseURL())
	}

	return c.JSON(&themes)
}

func GetTheme(c *fiber.Ctx) error {
	id, err := ParseUintParam(c, "id")
	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var theme entities.Theme
	config.Database.Find(&theme, "id = ?", id)
	if theme.Title == "" {
		return ErrorResponse(c, fiber.StatusBadRequest, "Theme not found!")
	}
	rewriteTheme(&theme, c.BaseURL())
	return c.JSON(theme)
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

	themeConf, err := c.FormFile("themeConfig")
	if err == nil {
		theme.ThemesConfig, err = SaveFile(c, themeConf)
		if err != nil {
			return ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
		}
	}

	// set approved false by default
	theme.Approved = false
	theme.Username = user.Name
	theme.Uploaded = time.Now().Unix()

	config.Database.Create(&theme)

	rewriteTheme(&theme, c.BaseURL())
	return c.Status(fiber.StatusCreated).JSON(theme)
}

func ApproveTheme(c *fiber.Ctx) error {
	user, err := ValidateUser(c)
	if err != nil {
		return nil
	}
	if user.Role == 0 {
		return ErrorResponse(c, fiber.StatusForbidden, "No permissions!")
	}
	id, err := ParseUintQuery(c, "id")
	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	config.Database.Model(&entities.Theme{}).Where("id = ?", uint(id)).Update("approved", true)
	return OkResponse(c)
}

func DeleteTheme(c *fiber.Ctx) error {
	user, err := ValidateUser(c)
	if err != nil {
		return nil
	}
	id, err := ParseUintQuery(c, "id")
	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var theme entities.Theme
	config.Database.Find(&theme, "id = ?", id)
	if isBlank(theme.Username) {
		return ErrorResponse(c, fiber.StatusBadRequest, "Theme not found!")
	}

	if user.Name != theme.Username && user.Role == 0 {
		return ErrorResponse(c, fiber.StatusForbidden, "No permissions to delete the theme!")
	}

	config.Database.Delete(&theme)
	return OkResponse(c)
}

func rewriteTheme(theme *entities.Theme, baseUrl string) {
	theme.Config = rewriteURL(baseUrl, theme.Config)
	theme.ThemesConfig = rewriteURL(baseUrl, theme.ThemesConfig)
	theme.Screenshot = rewriteURL(baseUrl, theme.Screenshot)
}

func rewriteURL(baseUrl string, path string) string {
	urlPath, _ := url.Parse(path)
	return baseUrl + urlPath.Path
}

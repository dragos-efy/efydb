package handlers

import (
	"strings"
	"time"

	"github.com/efydb/config"
	"github.com/efydb/entities"
	"github.com/efydb/util"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func GetThemes(c *fiber.Ctx) error {
	showUnapproved := c.Query("unapproved", "false") == "true"
	username := c.Query("username", "")

	var themes []entities.Theme

	query := config.Database
	if !showUnapproved {
		query = query.Where("approved = ?", true)
	}

	if username != "" {
		query = query.Where("username = ?", username)
	}

	query.Find(&themes)

	for index := range themes {
		rewriteTheme(&themes[index], c)
	}

	return c.JSON(&themes)
}

func GetTheme(c *fiber.Ctx) error {
	id, err := util.ParseUintParam(c, "id")
	if err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var theme entities.Theme
	config.Database.Find(&theme, "id = ?", id)
	if theme.Title == "" {
		return util.ErrorResponse(c, fiber.StatusBadRequest, "Theme not found!")
	}
	rewriteTheme(&theme, c)
	return c.JSON(theme)
}

func createFilesForTheme(c *fiber.Ctx) (entities.Theme, error) {
	var theme entities.Theme

	// get the uploaded screenshot
	ss, err := c.FormFile("screenshot")

	if err != nil {
		return theme, util.ErrorResponse(c, fiber.StatusBadRequest, "Screenshot missing!")
	}

	conf, err := c.FormFile("config")

	if err != nil {
		return theme, util.ErrorResponse(c, fiber.StatusBadRequest, "No config provided!")
	}

	data := c.FormValue("data", "")

	if data == "" {
		return theme, util.ErrorResponse(c, fiber.StatusBadRequest, "Data can't be empty!")
	}

	reader := strings.NewReader(data)
	jsonErr := json.NewDecoder(reader).Decode(&theme)

	if jsonErr != nil {
		return theme, util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if util.IsBlank(theme.Title) {
		return theme, util.ErrorResponse(c, fiber.StatusBadRequest, "Title can't be blank")
	}

	// create the new files and save them
	theme.Config, err = util.SaveFile(c, conf)
	if err != nil {
		return theme, util.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	theme.Screenshot, err = util.SaveFile(c, ss)
	if err != nil {
		return theme, util.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	themeConf, err := c.FormFile("imageConfig")
	if err == nil {
		theme.ImageConfig, err = util.SaveFile(c, themeConf)
		if err != nil {
			return theme, util.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
		}
	}

	return theme, nil
}

func CreateTheme(c *fiber.Ctx) error {
	// get the user
	user, err := util.ValidateUser(c)
	if err != nil {
		return nil
	}

	theme, err := createFilesForTheme(c)
	if err != nil {
		return nil
	}

	// set approved false by default
	theme.Approved = false
	theme.Username = user.Name
	theme.Uploaded = time.Now().Unix()

	config.Database.Create(&theme)

	rewriteTheme(&theme, c)
	return c.Status(fiber.StatusCreated).JSON(theme)
}

func EditTheme(c *fiber.Ctx) error {
	user, err := util.ValidateUser(c)
	if err != nil {
		return nil
	}

	id, err := util.ParseUintParam(c, "id")
	if err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var theme entities.Theme
	config.Database.Find(&theme, "id = ?", id)
	if theme.Username != user.Name {
		return util.ErrorResponse(c, fiber.StatusForbidden, "No permissions to edit the theme!")
	}

	newTheme, err := createFilesForTheme(c)
	if err != nil {
		return nil
	}

	theme.Title = newTheme.Title
	theme.Description = newTheme.Description
	theme.Config = newTheme.Config
	theme.EfyVersion = newTheme.EfyVersion
	theme.ImageConfig = newTheme.ImageConfig
	theme.Screenshot = newTheme.Screenshot
	theme.Tags = newTheme.Tags

	config.Database.Where("id = ?", theme.ID).Updates(theme)

	return c.JSON(theme)
}

func ApproveTheme(c *fiber.Ctx) error {
	user, err := util.ValidateUser(c)
	if err != nil {
		return nil
	}
	if user.Role == 0 {
		return util.ErrorResponse(c, fiber.StatusForbidden, "No permissions!")
	}
	id, err := util.ParseUintQuery(c, "id")
	if err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	config.Database.Model(&entities.Theme{}).Where("id = ?", uint(id)).Update("approved", true)
	return util.OkResponse(c)
}

func DeleteTheme(c *fiber.Ctx) error {
	user, err := util.ValidateUser(c)
	if err != nil {
		return nil
	}
	id, err := util.ParseUintQuery(c, "id")
	if err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var theme entities.Theme
	config.Database.Find(&theme, "id = ?", id)
	if util.IsBlank(theme.Username) {
		return util.ErrorResponse(c, fiber.StatusBadRequest, "Theme not found!")
	}

	if user.Name != theme.Username && user.Role == 0 {
		return util.ErrorResponse(c, fiber.StatusForbidden, "No permissions to delete the theme!")
	}

	config.Database.Delete(&theme)
	return util.OkResponse(c)
}

func rewriteTheme(theme *entities.Theme, c *fiber.Ctx) {
	baseUrl := c.BaseURL()
	theme.Config = baseUrl + theme.Config
	theme.ImageConfig = baseUrl + theme.ImageConfig
	theme.Screenshot = baseUrl + theme.Screenshot
}

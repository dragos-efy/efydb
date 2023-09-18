package handlers

import (
	"errors"
	"math"
	"strconv"
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
	sortOrder := c.Query("sort", "score")
	page, err := util.ParseUintQuery(c, "page")
	if err != nil {
		page = 1
	}
	limit, err := util.ParseUintQuery(c, "limit")
	if err != nil {
		limit = 20
	}
	searchQuery := c.Query("q")

	var themes []entities.Theme

	query := config.Database.Model(&themes)
	if !showUnapproved {
		query = query.Where("approved = ?", true)
	}

	if username != "" {
		query = query.Where("username = ?", username)
	}

	if !util.IsBlank(searchQuery) {
		q := "%" + searchQuery + "%"
		query = query.Where("title LIKE ? OR description LIKE ? OR username LIKE ?", q, q, q)
	}

	if sortOrder == "score" {
		query = query.Order("score desc")
	}

	if page != 1 {
		offset := (int(page) - 1) * int(limit)
		query = query.Offset(offset)
	}

	query.Limit(int(limit)).Find(&themes)

	for index := range themes {
		rewriteTheme(&themes[index], c)
	}

	return c.JSON(&themes)
}

func loadThemeById(id uint, c *fiber.Ctx) (entities.Theme, error) {
	var theme entities.Theme
	config.Database.Find(&theme, "id = ?", id)

	if theme.Title == "" {
		return theme, errors.New("Theme not found!")
	}

	if user, err := util.ValidateUser(c); err == nil {
		loadVoteByUser(&theme, user)
	}

	return theme, nil
}

func GetTheme(c *fiber.Ctx) error {
	id, err := util.ParseUintParam(c, "id")
	if err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	theme, err := loadThemeById(id, c)
	if err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
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

	id, err := strconv.ParseUint(c.Query("id"), 10, 32)
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

	deleteThemeFiles(theme)

	theme.Title = newTheme.Title
	theme.Description = newTheme.Description
	theme.Config = newTheme.Config
	theme.ImageConfig = newTheme.ImageConfig
	theme.Screenshot = newTheme.Screenshot

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

func deleteThemeFiles(theme entities.Theme) {
	util.DeleteFile(theme.Config)
	util.DeleteFile(theme.ImageConfig)
	util.DeleteFile(theme.Screenshot)
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

	deleteThemeFiles(theme)
	config.Database.Where(&theme).Delete(&theme)

	return util.OkResponse(c)
}

func VoteTheme(c *fiber.Ctx) error {
	user, err := util.ValidateUser(c)
	if err != nil {
		return nil
	}

	themeId, err := util.ParseUintParam(c, "id")
	if err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, "Theme not found!")
	}

	score, err := strconv.Atoi(c.Query("score"))
	if err != nil || math.Abs(float64(score)) > 1 {
		return util.ErrorResponse(c, fiber.StatusBadRequest, "Invalid score specified!")
	}

	// delete the old vote by this user if existent
	config.Database.Where(&entities.Vote{ThemeID: themeId, UserID: user.ID}).Delete(&entities.Vote{})

	if score != 0 {
		vote := entities.Vote{
			UserID:  user.ID,
			ThemeID: themeId,
			Score:   score,
		}
		config.Database.Create(&vote)
	}

	theme, _ := loadThemeById(themeId, c)
	theme.Score = getScore(theme)
	config.Database.Model(&theme).Where("id = ?", theme.ID).Update("score", theme.Score)

	rewriteTheme(&theme, c)
	return c.JSON(theme)
}

func rewriteTheme(theme *entities.Theme, c *fiber.Ctx) {
	baseUrl := c.BaseURL()
	theme.Config = rewriteUrl(baseUrl, theme.Config)
	theme.ImageConfig = rewriteUrl(baseUrl, theme.ImageConfig)
	theme.Screenshot = rewriteUrl(baseUrl, theme.Screenshot)
}

func rewriteUrl(baseUrl string, url string) string {
	if util.IsBlank(url) {
		return url
	}
	return baseUrl + url
}

func getScore(theme entities.Theme) int {
	var votes []entities.Vote
	config.Database.Where(&entities.Vote{ThemeID: theme.ID}).Find(&votes)

	score := 0
	for _, vote := range votes {
		score += vote.Score
	}
	return score
}

func loadVoteByUser(theme *entities.Theme, user entities.User) {
	vote := entities.Vote{
		UserID:  user.ID,
		ThemeID: theme.ID,
	}
	config.Database.Where(&vote).Find(&vote)
	theme.UserScore = vote.Score
}

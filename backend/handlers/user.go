package handlers

import (
	"github.com/efydb/config"
	"github.com/efydb/entities"
	"github.com/efydb/util"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(entities.User)

	if err := c.BodyParser(user); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if util.IsBlank(user.Name) || util.IsBlank(user.Password) {
		return util.ErrorResponse(c, fiber.StatusBadRequest, "Username and Password can't be empty!")
	}

	if len(user.Password) < 8 {
		return util.ErrorResponse(c, fiber.StatusBadRequest, "Password must be 8 chars at minimum!")
	}

	var nameTaken bool
	config.Database.Model(&entities.User{}).
		Select("count(*) > 0").
		Select("name = ?", user.Name).
		Find(&nameTaken)

	if nameTaken {
		return util.ErrorResponse(c, fiber.StatusForbidden, "Username already taken!")
	}

	user.Password, _ = util.HashPassword(user.Password)
	user.Token = util.GenerateSecureToken(20)

	// Grant owner permissions to the first registerd user
	if len(getAllUsers()) == 0 {
		user.Role = 2
	}

	config.Database.Create(&user)

	user.Password = ""
	return c.JSON(&user)
}

func GetUsers(c *fiber.Ctx) error {
	user, err := util.ValidateUser(c)
	if err != nil {
		return nil
	}

	if user.Role == 0 {
		return util.ErrorResponse(c, fiber.StatusForbidden, "Not privileged enough!")
	}

	users := getAllUsers()
	return c.JSON(&users)
}

func removeSecretUserInfo(user *entities.User) {
	user.Token = ""
	user.Password = ""
}

func getAllUsers() []entities.User {
	var users []entities.User
	config.Database.Find(&users)
	for index := range users {
		removeSecretUserInfo(&users[index])
	}
	return users
}

func DeleteUser(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	query := config.Database.Delete(&entities.User{}, "token = ?", token)
	if query.RowsAffected == 0 {
		return util.ErrorResponse(c, fiber.StatusBadRequest, "Invalid Access Token!")
	}
	return util.OkResponse(c)
}

func LoginUser(c *fiber.Ctx) error {
	userReq := new(entities.User)
	if err := c.BodyParser(userReq); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var user entities.User
	err := config.Database.Find(&user, "name = ?", userReq.Name).Error

	if err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if !util.CheckPasswordHash(userReq.Password, user.Password) {
		return util.ErrorResponse(c, fiber.StatusForbidden, "Password doesn't match!")
	}

	user.Password = ""
	return c.JSON(&user)
}

func PromoteUser(c *fiber.Ctx) error {
	user, err := util.ValidateUser(c)

	if err != nil {
		return err
	}

	if user.Role == 0 {
		return util.ErrorResponse(c, fiber.StatusForbidden, "No permissions for promoting!")
	}

	userToChange := new(entities.User)
	if err := c.BodyParser(userToChange); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if userToChange.Role < 0 || userToChange.Role > 2 {
		return util.ErrorResponse(c, fiber.StatusBadRequest, "Invalid new role!")
	}

	if userToChange.Role > user.Role {
		return util.ErrorResponse(c, fiber.StatusForbidden, "Not privileged enough!")
	}

	query := config.Database.Find(&user, user)
	if query.Error != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	user.Role = userToChange.Role
	config.Database.Where("id = ?", user.ID).Updates(userToChange)

	return util.OkResponse(c)
}

func UpdateUser(c *fiber.Ctx) error {
	user, err := util.ValidateUser(c)
	if err != nil {
		return nil
	}

	updatedUser := new(entities.User)
	if err := c.BodyParser(updatedUser); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if !util.IsBlank(updatedUser.Password) {
		user.Password, _ = util.HashPassword(updatedUser.Password)
	}

	if !util.IsBlank(updatedUser.Bio) {
		user.Bio = updatedUser.Bio
	}

	config.Database.Updates(&user)

	return util.OkResponse(c)
}

func GetUser(c *fiber.Ctx) error {
	user, err := util.ValidateUser(c)
	if err != nil {
		return nil
	}
	return c.JSON(user)
}

func GetUserInfo(c *fiber.Ctx) error {
	name := c.Params("name")

	if util.IsBlank(name) {
		return util.ErrorResponse(c, fiber.StatusBadRequest, "No valid username provided!")
	}

	var user entities.User
	if config.Database.Where("name = ?", name).Find(&user).RowsAffected == 0 {
		return util.ErrorResponse(c, fiber.StatusBadRequest, "User not found!")
	}
	removeSecretUserInfo(&user)

	var themes []entities.Theme
	config.Database.Where("username = ?", name).Find(&themes)
	for index := range themes {
		rewriteTheme(&themes[index], c)
	}

	userInfo := entities.UserInfo{
		User:   user,
		Themes: themes,
	}
	return c.JSON(userInfo)
}

package handlers

import (
	"github.com/efydb/config"
	"github.com/efydb/entities"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(entities.User)

	if err := c.BodyParser(user); err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if isBlank(user.Name) || isBlank(user.Password) {
		return ErrorResponse(c, fiber.StatusBadRequest, "Username and Password can't be empty!")
	}

	if len(user.Password) < 8 {
		return ErrorResponse(c, fiber.StatusBadRequest, "Password must be 8 chars at minimum!")
	}

	var nameTaken bool
	config.Database.Model(&entities.User{}).
		Select("count(*) > 0").
		Select("name = ?", user.Name).
		Find(&nameTaken)

	if nameTaken {
		return ErrorResponse(c, fiber.StatusForbidden, "Username already taken!")
	}

	user.Password, _ = HashPassword(user.Password)
	user.Token = GenerateSecureToken(20)

	// Grant owner permissions to the first registerd user
	if len(getAllUsers()) == 0 {
		user.Role = 2
	}

	config.Database.Create(&user)

	user.Password = ""
	return c.JSON(&user)
}

func GetUsers(c *fiber.Ctx) error {
	users := getAllUsers()
	return c.JSON(&users)
}

func getAllUsers() []entities.User {
	var users []entities.User
	config.Database.Find(&users)
	for index := range users {
		users[index].Token = ""
		users[index].Password = ""
	}
	return users
}

func DeleteUser(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	query := config.Database.Delete(&entities.User{}, "token = ?", token)
	if query.RowsAffected == 0 {
		return ErrorResponse(c, fiber.StatusBadRequest, "Invalid Access Token!")
	}
	return OkResponse(c)
}

func LoginUser(c *fiber.Ctx) error {
	userReq := new(entities.User)
	if err := c.BodyParser(userReq); err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var user entities.User
	err := config.Database.Find(&user, "name = ?", userReq.Name).Error

	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if !CheckPasswordHash(userReq.Password, user.Password) {
		return ErrorResponse(c, fiber.StatusForbidden, "Password doesn't match!")
	}

	user.Password = ""
	return c.JSON(&user)
}

func PromoteUser(c *fiber.Ctx) error {
	user, err := ValidateUser(c)

	if err != nil {
		return err
	}

	if user.Role == 0 {
		return ErrorResponse(c, fiber.StatusForbidden, "No permissions for promoting!")
	}

	userToChange := new(entities.User)
	if err := c.BodyParser(userToChange); err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if userToChange.Role < 0 || userToChange.Role > 2 {
		return ErrorResponse(c, fiber.StatusBadRequest, "Invalid new role!")
	}

	if userToChange.Role > user.Role {
		return ErrorResponse(c, fiber.StatusForbidden, "Not priviliged enough!")
	}

	query := config.Database.Find(&user, user)
	if query.Error != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	user.Role = userToChange.Role
	config.Database.Where("id = ?", user.ID).Updates(userToChange)

	return OkResponse(c)
}

func UpdateUser(c *fiber.Ctx) error {
	user, err := ValidateUser(c)
	if err != nil {
		return nil
	}

	updatedUser := new(entities.User)
	if err := c.BodyParser(updatedUser); err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if !isBlank(updatedUser.Password) {
		user.Password, _ = HashPassword(updatedUser.Password)
	}

	if !isBlank(updatedUser.Bio) {
		user.Bio = updatedUser.Bio
	}

	config.Database.Updates(&user)

	return OkResponse(c)
}

func GetUser(c *fiber.Ctx) error {
	user, err := ValidateUser(c)
	if err != nil {
		return nil
	}
	return c.JSON(user)
}

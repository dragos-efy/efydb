package handlers

import (
	"github.com/efydb/config"
	"github.com/efydb/entities"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(entities.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": err.Error(),
			},
		)
	}

	if isBlank(user.Name) || isBlank(user.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(
			entities.Message{
				Message: "Username and Password can't be empty!",
			},
		)
	}

	if len(user.Password) < 8 {
		return c.Status(fiber.StatusBadRequest).JSON(
			entities.Message{
				Message: "Password must be 8 chars at minimum!",
			},
		)
	}

	var nameTaken bool
	config.Database.Model(&entities.User{}).
		Select("count(*) > 0").
		Select("name = ?", user.Name).
		Find(&nameTaken)

	if nameTaken {
		return c.Status(fiber.StatusForbidden).JSON(
			entities.Message{
				Message: "Username already taken!",
			},
		)
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
	token := c.GetReqHeaders()["Authorization"]
	query := config.Database.Delete(&entities.User{}, "token = ?", token)
	if query.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(
			entities.Message{
				Message: "Invalid access token!",
			},
		)
	}
	return c.JSON(
		entities.Message{
			Message: "ok",
		},
	)
}

func LoginUser(c *fiber.Ctx) error {
	user, err := ValidateUser(c)
	if err != nil {
		return err
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
		return c.Status(fiber.StatusForbidden).JSON(
			entities.Message{
				Message: "No permissions for promoting!",
			},
		)
	}

	userToPromote := new(entities.User)
	if err := c.BodyParser(userToPromote); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			entities.Message{
				Message: err.Error(),
			},
		)
	}

	query := config.Database.Find(&user, user)
	if query.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			entities.Message{
				Message: err.Error(),
			},
		)
	}

	user.Role = user.Role + 1
	config.Database.Where("id = ?", user.ID).Updates(userToPromote)

	return c.Status(200).JSON(
		entities.Message{
			Message: "ok",
		},
	)
}

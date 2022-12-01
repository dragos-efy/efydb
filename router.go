package main

import (
	"log"

	"github.com/efydb/handlers"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CreateRouter() {
	router := fiber.New(
		fiber.Config{
			BodyLimit:   20 * 1024 * 1024,
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		},
	)

	router.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	router.Get("/", func(c *fiber.Ctx) error {
		return handlers.OkResponse(c)
	})

	router.Static("/files", "./files")

	user := router.Group("/users")
	user.Get("/", handlers.GetUsers)
	user.Post("/register", handlers.CreateUser)
	user.Post("/login", handlers.LoginUser)
	user.Patch("/update", handlers.UpdateUser)
	user.Post("/promote", handlers.PromoteUser)
	user.Delete("/delete", handlers.DeleteUser)

	themes := router.Group("/themes")
	themes.Get("/", handlers.GetThemes)
	themes.Post("/create", handlers.CreateTheme)
	themes.Delete("/delete", handlers.DeleteTheme)
	themes.Post("/approve", handlers.ApproveTheme)

	log.Fatal(router.Listen(":8000"))
}

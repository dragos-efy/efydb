package main

import (
	"log"

	"github.com/efydb/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CreateRouter() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	router.Post("/user/register", handlers.CreateUser)
	router.Post("/user/login", handlers.LoginUser)
	router.Post("/user/promote", handlers.PromoteUser)
	router.Delete("/user/delete", handlers.DeleteUser)
	router.Get("/users", handlers.GetUsers)

	log.Fatal(router.Listen(":8000"))
}

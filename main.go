package main

import (
	"github.com/gofiber/fiber/v2"
	"rest-api-gorm-fiber/config"
	"rest-api-gorm-fiber/handlers"
)

func main() {
	var err error
	app := fiber.New()

	err = config.Connect()
	if err != nil {
		return
	}

	app.Get("/user", handlers.GetUser)
	app.Post("/user", handlers.AddUser)
	app.Get("/user/:id", handlers.GetUserById)
	app.Delete("/user/:id", handlers.DeleteUserById)
	app.Put("/user/:id", handlers.UpdateUser)

	err = app.Listen(": 3000")
	if err != nil {
		return
	}
}

package main

import (
	"log"
	"github.com/shanliao420/learn1000/back-go/router"
	"github.com/gofiber/fiber/v3"
)


func main() {
    app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})


	contentRouter := &router.ContentRouter{}

	api := app.Group("api")

	contentRouter.RegisterContentRouter(api)

    log.Fatal(app.Listen(":3000"))
}
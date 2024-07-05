package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/shanliao420/learn1000/back-go/content"
	"github.com/shanliao420/learn1000/back-go/router"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/api/save", func(c fiber.Ctx) error {
		id := c.FormValue("id")
		root := content.GetItemRoot(id)
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		return c.SaveFile(file, root+file.Filename)
	})

	contentRouter := router.NewContentRouter()

	api := app.Group("api")

	contentRouter.RegisterContentRouter(api)

	log.Fatal(app.Listen(":3000"))
}

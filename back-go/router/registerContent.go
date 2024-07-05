package router

import (
	"github.com/gofiber/fiber/v3/log"

	"github.com/gofiber/fiber/v3"
	"github.com/shanliao420/learn1000/back-go/do"
	"github.com/shanliao420/learn1000/back-go/service"
)

type ContentRouter struct {
	ContentService *service.ContentService
}

func NewContentRouter() *ContentRouter {
	return &ContentRouter{
		ContentService: service.NewContentService(),
	}
}

func (c *ContentRouter) RegisterContentRouter(router fiber.Router) {
	router.Get("/list", func(ctx fiber.Ctx) error {
		ctx.JSON(c.ContentService.List())
		return nil
	})

	router.Get("/get/:id", func(ctx fiber.Ctx) error {
		ctx.JSON(c.ContentService.Get(ctx.Params("id")))
		return nil
	})

	router.Post("/update/:id", func(ctx fiber.Ctx) error {
		ctx.FormValue("source")
		ctx.FormValue("translation")
		item := &do.Item{
			Source:      ctx.FormValue("source"),
			Translation: ctx.FormValue("translation"),
			Id:          ctx.Params("id"),
		}
		c.ContentService.Update(*item)
		return nil
	})

	router.All("/create", func(ctx fiber.Ctx) error {
		c.ContentService.NewItem()
		log.Info("create")
		return nil
	})

}

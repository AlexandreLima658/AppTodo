package router

import (
	"github.com/AlexandreLima658/AppTodo/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Post("/", func(ctx *fiber.Ctx) error {
		return controllers.CreateTodo(ctx)
	})
	app.Get("/", func(ctx *fiber.Ctx) error {
		return controllers.SearchTodos(ctx)
	})
	app.Put("/update/:id", func(ctx *fiber.Ctx) error {
		return controllers.UpdatedTodo(ctx)
	})
	app.Delete("/delete/:id", func(ctx *fiber.Ctx) error {
		return controllers.DeleteTodo(ctx)
	})
}

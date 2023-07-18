package main

import (
	"github.com/AlexandreLima658/AppTodo/src/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
)

func main() {
	engine := html.New("./src/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	router.Router(app)
	app.Static("/", "./src/views")
	log.Fatal(app.Listen(":8080"))
}

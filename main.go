package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

func main() {
	engine := django.New("./views", ".django")

	app := fiber.New(fiber.Config{
		Views:         engine,
		Prefork:       true,
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  "Fiber",
		AppName:       "Test App v1.0.1",
	})

	app.Static("/", "./public")

	app.Get("/api/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("200 ok")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello World",
		}, "layouts/main")
	})

	app.Listen(":8080")
}

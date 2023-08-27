package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  "Fiber",
		AppName:       "Test App v1.0.1",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Haha")
	})

	app.Listen(":8080")
}

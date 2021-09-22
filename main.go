package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// app.Get("/:name", func(c *fiber.Ctx) error {
	// 	return c.SendString(fmt.Sprintf("Hello, %s", c.Params("name")))
	// })

	// TODO: implement handler!
	app.Get("/delete/limit/:key/*", ResetLimitHandler)

	app.Get("/delete/cache/:key/*", EvictCacheHandler)

	app.Get("/:key/*", ProxyHandler)

	app.Listen(":3000")
}

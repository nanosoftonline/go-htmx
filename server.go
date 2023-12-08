package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func htmxRender(page string, data fiber.Map, c *fiber.Ctx) error {

	if c.Get("hx-request") == "true" {
		return c.Render(page, data, "")
	} else {
		return c.Render(page, data, "layouts/main")
	}

}

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(
		fiber.Config{
			Views:       engine,
			ViewsLayout: "layouts/main",
		},
	)

	app.Get("/", func(c *fiber.Ctx) error {

		return htmxRender("home", fiber.Map{}, c)
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		return htmxRender("users", fiber.Map{}, c)
	})

	app.Get("/posts", func(c *fiber.Ctx) error {
		return htmxRender("posts", fiber.Map{}, c)
	})

	app.Use(func(c *fiber.Ctx) error {
		return htmxRender("404", fiber.Map{}, c)
	})

	log.Fatal(app.Listen(":3001"))
}

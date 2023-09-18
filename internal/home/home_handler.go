package home

import (
	"github.com/bitstorm-tech/cockaigne/internal/auth"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		if !auth.IsAuthenticated(c) {
			return c.Redirect("/login")
		}
		return c.Render("pages/home", nil, "layouts/main")
	})
}

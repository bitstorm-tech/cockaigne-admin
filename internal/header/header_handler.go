package header

import (
	"github.com/bitstorm-tech/cockaigne/internal/auth"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	app.Get("/ui/header", func(c *fiber.Ctx) error {
		showMenu := c.Query("showMenu")
		isAuthenticated := auth.IsAuthenticated(c)

		return c.Render("partials/header", fiber.Map{
			"showMenu":        showMenu == "true",
			"isAuthenticated": isAuthenticated,
		})
	})

	app.Get("/partials/alert", func(c *fiber.Ctx) error {
		return c.Render("partials/alert", fiber.Map{
			"message": "Läuft doch eigentlich ganz gut, oder?",
		})
	})

	app.Delete("/ui/remove", func(c *fiber.Ctx) error {
		return c.SendString("")
	})
}

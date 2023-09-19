package voucher

import (
	"github.com/bitstorm-tech/cockaigne/internal/auth"
	"github.com/bitstorm-tech/cockaigne/internal/persistence"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func Register(app *fiber.App) {
	app.Get("/vouchers", func(c *fiber.Ctx) error {
		if !auth.IsAuthenticated(c) {
			return c.Redirect("/login")
		}

		return c.Render("pages/vouchers", nil, "layouts/main")
	})

	app.Get("/create-voucher", func(c *fiber.Ctx) error {
		if !auth.IsAuthenticated(c) {
			return c.Redirect("/login")
		}

		return c.Render("pages/create-voucher", nil, "layouts/main")
	})

	app.Post("/api/vouchers", func(c *fiber.Ctx) error {
		if !auth.IsAuthenticated(c) {
			return c.Redirect("/login")
		}

		voucher := &Voucher{}
		err := c.BodyParser(voucher)

		if err != nil {
			log.Warnf("Can't parse voucher from request body: %+v", err)
		}

		err = persistence.DB.Create(voucher).Error

		if err != nil {
			log.Warnf("Can't create voucher: %+v", err)
		}

		return c.Redirect("/vouchers")
	})
}

package voucher

import (
	"github.com/bitstorm-tech/cockaigne/internal/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func Register(app *fiber.App) {
	app.Get("/vouchers", func(c *fiber.Ctx) error {
		jwt := c.Cookies("jwt")
		_, err := auth.ParseJwtToken(jwt)

		if err != nil {
			log.Errorf("Can't parse JWT token: %+v", err)
			return c.Redirect("/login")
		}

		return c.Render("pages/vouchers", nil, "layouts/main")
	})

	app.Post("/api/vouchers", func(c *fiber.Ctx) error {
		jwt := c.Cookies("jwt")
		_, err := auth.ParseJwtToken(jwt)

		if err != nil {
			log.Errorf("Can't parse JWT token: %+v", err)
			return c.Redirect("/login")
		}

		voucher := &Voucher{}
		err = c.BodyParser(voucher)

		if err != nil {
			log.Warnf("Can't save new voucher: %+v", err)
		}

		log.Debugf("New voucher: %+v", voucher)

		return c.Render("pages/vouchers", nil, "layouts/main")
	})
}

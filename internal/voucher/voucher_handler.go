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

		vouchers := []Voucher{}
		persistence.DB.Find(&vouchers)

		return c.Render("pages/vouchers", fiber.Map{"vouchers": vouchers}, "layouts/main")
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

		voucherRequest := CreateVoucherRequest{}
		err := c.BodyParser(&voucherRequest)
		if err != nil {
			log.Warnf("Can't parse voucher from request body: %+v", err)
		}

		voucher, err := voucherRequest.ToVoucher()
		if err != nil {
			log.Warnf("Can't create voucher from request: %+v", err)
		}

		err = persistence.DB.Create(&voucher).Error

		if err != nil {
			log.Warnf("Can't create voucher in DB: %+v", err)
		}

		return c.Redirect("/vouchers")
	})
}

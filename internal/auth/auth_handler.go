package auth

import (
	"github.com/bitstorm-tech/cockaigne/internal/persistence"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
)

func Register(app *fiber.App) {
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("pages/login", nil, "layouts/main")
	})

	app.Post("/api/login", login)

	app.Get("/api/logout", logout)
}

func login(c *fiber.Ctx) error {
	request := LoginRequest{}
	err := c.BodyParser(&request)

	if err != nil {
		log.Errorf("Error while signup %v", err)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	log.Debugf("Login attempt: %+v", request.Username)

	acc := AdminAccount{}
	err = persistence.DB.Where("username ilike ?", request.Username).First(&acc).Error

	if err != nil {
		return c.Render("partials/alert", fiber.Map{"message": "Benutzername oder Passwort falsch"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(request.Password))

	if err != nil {
		return c.Render("partials/alert", fiber.Map{"message": "Benutzername oder Passwort falsch"})
	}

	c.Set("HX-Location", "/")

	jwt := CreateJwtToken(acc)
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    jwt,
		HTTPOnly: true,
	})

	return nil
}

func logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		HTTPOnly: true,
	})

	return c.Redirect("/login")
}

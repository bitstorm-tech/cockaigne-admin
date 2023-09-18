package auth

import (
	"github.com/gofiber/fiber/v2"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func CreateJwtToken(acc AdminAccount) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": acc.ID,
	})

	signedString, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Errorf("Can't signe JWT token: %+v", err)
	}

	return signedString
}

func ParseJwtToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}

func IsAuthenticated(c *fiber.Ctx) bool {
	_, err := ParseJwtToken(c.Cookies("jwt"))
	return err == nil
}

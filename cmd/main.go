package main

import (
	"github.com/bitstorm-tech/cockaigne/internal/home"
	"log"
	"os"
	"strings"

	"github.com/bitstorm-tech/cockaigne/internal/auth"
	"github.com/bitstorm-tech/cockaigne/internal/header"
	"github.com/bitstorm-tech/cockaigne/internal/persistence"
	"github.com/bitstorm-tech/cockaigne/internal/voucher"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	hostAndPort := host + ":" + port

	log.Printf("Starting Cockaigne server (on %s) ...", hostAndPort)

	persistence.ConnectToDb()

	migrateDb := strings.ToLower(os.Getenv("MIGRATE_DATABASE")) == "true"

	if migrateDb {
		err := persistence.DB.AutoMigrate()
		if err != nil {
			log.Fatal(err)
		}
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./static")

	auth.Register(app)
	home.Register(app)
	header.Register(app)
	voucher.Register(app)

	log.Fatal(app.Listen(hostAndPort))
}

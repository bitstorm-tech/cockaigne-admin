package main

import (
	"flag"
	"fmt"
	"github.com/bitstorm-tech/cockaigne/internal/auth"
	"github.com/bitstorm-tech/cockaigne/internal/persistence"
	"github.com/bitstorm-tech/cockaigne/internal/voucher"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

func main() {
	migrate := flag.Bool("m", false, "Migrate database")
	createAdminAccount := flag.String("c", "", "Create an admin account with a necessary comment")
	flag.Parse()

	println("Arguments: ", len(flag.Args()))

	if *migrate {
		println("Migrating database ...")
		persistence.ConnectToDb()
		err := persistence.DB.AutoMigrate(&auth.AdminAccount{}, &voucher.Voucher{})
		if err != nil {
			panic(err)
		}
	}

	if len(*createAdminAccount) > 0 {
		println(fmt.Sprintf("Creating admin account with comment: '%s' ...", *createAdminAccount))
		persistence.ConnectToDb()
		username := randomString(8)
		password := randomString(16)
		passwordEnc, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}

		account := auth.AdminAccount{
			Username: username,
			Password: string(passwordEnc),
			Comment:  *createAdminAccount,
			IsActive: true,
		}

		if err := persistence.DB.Create(&account).Error; err != nil {
			panic(err)
		}

		println("Admin account created:")
		fmt.Printf("Username: %s\n", username)
		fmt.Printf("Password: %s\n", password)
	}

	if !*migrate && len(*createAdminAccount) == 0 {
		flag.PrintDefaults()
	}
}

func randomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

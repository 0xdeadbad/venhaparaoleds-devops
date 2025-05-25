package main

import (
	"context"
	"fmt"
	"log"
	"os"
	_ "time"

	_ "github.com/0xdeadbad/venhaparaoleds-devops/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func appGoroutine(ctx context.Context, cancel context.CancelCauseFunc, app *fiber.App) {
	err := app.Listen(":3000")
	cancel(err)
}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// func main() {
// 	app := fiber.New()

// 	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
// 	defer cancel()

// 	app.Get("/swagger/*", swagger.HandlerDefault) // default

// 	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
// 		URL:         "http://example.com/doc.json",
// 		DeepLinking: false,
// 		// Expand ("list") or Collapse ("none") tag groups by default
// 		DocExpansion: "none",
// 		// Prefill OAuth ClientId on Authorize popup
// 		OAuth: &swagger.OAuthConfig{
// 			AppName:  "OAuth Provider",
// 			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
// 		},
// 		// Ability to change OAuth2 redirect uri location
// 		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
// 	}))

// 	done := make(chan any, 1)
// 	go func() {
// 		if err := app.Listen(":3000"); err != nil {
// 			fmt.Printf("%+v\n", err)
// 		}
// 		done <- struct{}{}
// 	}()

// main_loop:
//
//		for {
//			select {
//			case <-ctx.Done():
//				if err := app.Shutdown(); err != nil {
//					fmt.Printf("%+v\n", err)
//				}
//				break
//			case <-done:
//				break main_loop
//			}
//		}
//	}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file couldn't be loaded.")
	}
}

func main() {
	POSTGRES_HOSTNAME := os.Getenv("POSTGRES_HOSTNAME")
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_PORT := os.Getenv("POSTGRES_PORT")
	POSTGRES_SSL := os.Getenv("POSTGRES_SSL")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", POSTGRES_HOSTNAME, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB, POSTGRES_PORT, POSTGRES_SSL)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Applicant{})
}

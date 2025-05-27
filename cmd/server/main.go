package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	_ "time"

	_ "github.com/0xdeadbad/venhaparaoleds-devops/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgsqlCrudApi struct {
	db *gorm.DB
}

func (p *PgsqlCrudApi) Create(a ...any) (*Concourse, error) {

	if len(a) <= 0 {
		return nil, errors.New("Method Create(...) needs at least one argument")
	}

	if _, ok := a[0].(string); !ok {
		return nil, errors.New("Fist parameter of method Create(...) should be a string")
	}
	// TODO: write more checks

	org := a[0].(string)
	edital := a[1].(string)
	code := a[2].(uint64)

	obj := &Concourse{
		Org:       org,
		Edital:    edital,
		ConcCode:  code,
		Vacancies: []Vacancy{},
		Model:     gorm.Model{},
	}

	res := p.db.Create(obj)

	if res.Error != nil {
		return nil, res.Error
	}

	return obj, nil
}

func (p *PgsqlCrudApi) Read(a ...any) (*Concourse, error) {

	obj := &Concourse{ConcCode: a[0].(*Concourse).ConcCode}

	res := p.db.First(obj)
	if res.Error != nil {
		return nil, res.Error
	}

	return obj, nil
}
func (p *PgsqlCrudApi) Update(a ...any) (*Concourse, error) {

	return &Concourse{}, nil
}

func (p *PgsqlCrudApi) Delete(a ...any) (*Concourse, error) {

	fmt.Println("DELETE")

	return &Concourse{}, nil
}

func appGoroutine(ctx context.Context, cancel context.CancelCauseFunc, app *fiber.App) {
	err := app.Listen(":3000")
	cancel(err)
}

// Load .env as environment variables if possible
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file couldn't be loaded.")
	}
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
func main() {
	var POSTGRES_HOSTNAME, POSTGRES_PASSWORD, POSTGRES_USER, POSTGRES_PORT, POSTGRES_SSL, POSTGRES_DB string

	app := fiber.New()

	// api.AddApiGroup(app, crud)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	ok := false
	if POSTGRES_HOSTNAME, ok = os.LookupEnv("POSTGRES_HOSTNAME"); !ok {
		log.Fatalln("POSTGRE_HOSTNAME env variable not found")
	}
	if POSTGRES_PASSWORD, ok = os.LookupEnv("POSTGRES_PASSWORD"); !ok {
		log.Fatalln("POSTGRES_PASSWORD env variable not set")
	}
	if POSTGRES_USER, ok = os.LookupEnv("POSTGRES_USER"); !ok {
		log.Fatalln("POSTGRES_USER env variable not set")
	}
	if POSTGRES_PORT, ok = os.LookupEnv("POSTGRES_PORT"); !ok {
		log.Println("POSTGRES_PORT env variable not set. Using default port: 5432")
	}
	if POSTGRES_SSL, ok = os.LookupEnv("POSTGRES_SSL"); !ok {
		log.Println("POSTGRES_SSL env variable not set. Using default value: false")
	}
	if POSTGRES_DB, ok = os.LookupEnv("POSTGRES_DB"); !ok {
		log.Fatalln("POSTGRES_DB env variable not set")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", POSTGRES_HOSTNAME, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB, POSTGRES_PORT, POSTGRES_SSL)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&Applicant{}, &Concourse{}, &Profession{}, &Vacancy{})
	if err != nil {
		log.Fatalln(err)
	}

	crud := &PgsqlCrudApi{
		db: db,
	}

	obj, err := crud.Read(&Concourse{ConcCode: uint64(65433)})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v\n", obj)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	done := make(chan any, 1)
	go func() {
		if err := app.Listen(":3000"); err != nil {
			fmt.Printf("%+v\n", err)
		}
		done <- struct{}{}
	}()

main_loop:

	for {
		select {
		case <-ctx.Done():
			if err := app.Shutdown(); err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}
			break
		case <-done:
			break main_loop
		}
	}
}

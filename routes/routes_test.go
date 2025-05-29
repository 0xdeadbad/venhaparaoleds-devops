package routes

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	// "net/http"
	"testing"

	"github.com/0xdeadbad/venhaparaoleds-devops/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestApplicant(t *testing.T) {
	app := fiber.New()

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	err = db.AutoMigrate(&models.Applicant{}, &models.Concourse{}, &models.Profession{}, &models.Vacancy{})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	// ch := make(chan any)
	go func() {
		if err := app.Listen("0.0.0.0:3030"); err != nil {
			t.Fail()
		}
		// ch <- struct{}{}
	}()

	apiRoute := app.Group("/api")
	v1 := apiRoute.Group("/v1", func(c *fiber.Ctx) error { // middleware for /api/v1
		c.Set("Version", "v1")

		return c.Next()
	})

	if err = MainRouter(v1, db); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	g := &models.Applicant{}

	t.Run("Test Applicant POST", func(t *testing.T) {
		profs := []models.Profession{
			{
				Name: "test1",
			},
			{
				Name: "test2",
			},
		}

		g = &models.Applicant{
			Name:        "Garicas",
			Birth:       time.Now(),
			Email:       "tooowwwaaaaa",
			CPF:         "11121311131",
			Professions: profs,
		}

		// db.Model(g).Association("Professions").Append(&profs)

		data, err := json.Marshal(g)
		if err != nil {
			t.Fatalf("%s\n", err.Error())
		}

		b := bytes.NewReader(data)
		req, err := http.NewRequest("POST", "/api/v1/applicant", b)
		if err != nil {
			t.Fatalf("%s\n", err.Error())
		}

		req.Header.Set("Content-Type", "application/json")
		res, err := app.Test(req, -1)
		assert.Nil(t, err)
		assert.Equal(t, 200, res.StatusCode)
	})

	t.Run("Test Applicant GET", func(t *testing.T) {
		obj := new(models.Applicant)
		req, err := http.NewRequest("GET", "/api/v1/applicant/11121311131", nil)
		if err != nil {
			t.Fatalf("%s\n", err.Error())
		}

		res, err := app.Test(req, -1)
		assert.Nil(t, err)
		assert.Equal(t, 200, res.StatusCode)

		data, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("%s\n", err.Error())
		}
		defer res.Body.Close()

		m := fiber.Map{}
		err = json.Unmarshal(data, &m)
		if err != nil {
			t.Fatalf("%s\n", err.Error())
		}

		data2, err := json.Marshal(m["data"])
		if err != nil {
			t.Fatalf("%s\n", err.Error())
		}

		t.Logf("+++++++++++%s\n", string(data2))

		err = json.Unmarshal(data2, obj)
		if err != nil {
			t.Fatalf("%s\n", err.Error())
		}

		t.Logf("-------------%+v\n", obj)
	})

	// t.Run("Test Applicant PUT", func(t *testing.T) {
	// 	obj := &models.Applicant{
	// 		Name:  "Garicas",
	// 		Birth: time.Now(),
	// 		Email: "tooowwwaaaaa",
	// 		CPF:   "11121311131",
	// 		Professions: []models.Profession{
	// 			{
	// 				Name: "test1",
	// 			},
	// 			{
	// 				Name: "test2",
	// 			},
	// 		},
	// 	}
	// })
	defer app.Shutdown()
	// <-ch
}

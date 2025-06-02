package routes

import (
	"encoding/json"

	"github.com/0xdeadbad/venhaparaoleds-devops/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func MainRouter(r fiber.Router, db *gorm.DB) error {
	prof := r.Group("/profession")

	appl := r.Group("/applicant", func(c *fiber.Ctx) error {
		c.Accepts("json", "text")
		c.Accepts("application/json")

		return c.Next()
	})

	conc := r.Group("/concourse", func(c *fiber.Ctx) error {
		c.Accepts("json", "text")
		c.Accepts("application/json")

		return c.Next()
	})

	vac := r.Group("/vacancy")

	if err := applicantRouter(appl, db); err != nil {
		return err
	}

	if err := concourseRouter(conc, db); err != nil {
		return err
	}

	if err := vacancyRouter(vac, db); err != nil {
		return err
	}

	if err := professionRouter(prof, db); err != nil {
		return err
	}

	return nil
}

func applicantRouter(applicant fiber.Router, db *gorm.DB) error {
	applicant.Get("/:cpf", func(c *fiber.Ctx) error {
		var cpf string
		// objs := new([]models.Profession)
		obj := new(models.Applicant)

		if cpf = c.Params("cpf"); cpf == "" {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": "invalid cpf parameter",
			})
		}
		obj.CPF = cpf

		res := db.Where("cpf = ?", cpf).Preload("Professions").First(obj)
		if err := res.Error; err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}

		// var profs []models.Profession
		// err := db.Model(obj).Related()

		// log.Printf("%+v\n", obj.Professions)

		if count := res.RowsAffected; count <= 0 {
			return c.Status(404).JSON(fiber.Map{
				"status":  "error",
				"message": "record not found",
			})
		}

		data, err := json.Marshal(obj)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}

		m := fiber.Map{}
		if err := json.Unmarshal(data, &m); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"status": "success",
			"data":   m,
		})
	})

	applicant.Post("/", func(c *fiber.Ctx) error {
		obj := new(models.Applicant)

		if err := c.BodyParser(obj); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}

		res := db.Create(obj)
		if res.Error != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": res.Error.Error(),
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"status":   "success",
			"affected": res.RowsAffected,
		})
	})

	applicant.Put("/:cpf", func(c *fiber.Ctx) error {
		var cpf string
		obj := &models.Applicant{}

		if cpf = c.Params("cpf"); cpf == "" {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": "invalid cpf parameter",
			})
		}
		obj.CPF = cpf

		if err := c.BodyParser(obj); err != nil {
			return err
		}

		res := db.UpdateColumns(obj)
		if res.Error != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": res.Error.Error(),
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"status": "success",
		})
	})

	applicant.Delete("/:cpf", func(c *fiber.Ctx) error {
		var cpf string
		// obj := new(models.Applicant)

		if cpf = c.Params("cpf"); cpf == "" {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": "invalid cpf parameter",
			})
		}
		// obj.CPF = cpf

		res := db.Where("cpf = ?", cpf).Delete(&models.Applicant{})
		if res.Error != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": res.Error.Error(),
			})
		}

		if res.RowsAffected <= 0 {
			return c.Status(404).JSON(fiber.Map{
				"status":  "error",
				"message": "record not found",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"status":   "success",
			"affected": res.RowsAffected,
		})
	})

	return nil
}

func concourseRouter(concourse fiber.Router, db *gorm.DB) error {
	concourse.Get("/:code", func(c *fiber.Ctx) error {
		var code string
		obj := new(models.Concourse)

		if code = c.Params("code"); code == "" {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": "invalid code parameter",
			})
		}
		obj.ConcCode = code

		res := db.First(obj)
		if err := res.Error; err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}

		data, err := json.Marshal(obj)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}

		m := fiber.Map{}
		if err := json.Unmarshal(data, &m); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"status": "success",
			"data":   m,
		})
	})

	concourse.Post("/", func(c *fiber.Ctx) error {
		obj := new(models.Concourse)

		if err := c.BodyParser(obj); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}

		res := db.Create(obj)
		if res.Error != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": res.Error.Error(),
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"status": "success",
			"data":   obj,
		})
	})

	concourse.Delete("/:code", func(c *fiber.Ctx) error {
		var code string

		if code = c.Params("code"); code == "" {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": "invalid code parameter",
			})
		}

		res := db.Where("conc_code = ?", code).Delete(&models.Concourse{})
		if res.Error != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": res.Error.Error(),
			})
		}

		if res.RowsAffected <= 0 {
			return c.Status(404).JSON(fiber.Map{
				"status":   "error",
				"affected": res.RowsAffected,
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"status":   "success",
			"affected": res.RowsAffected,
		})
	})

	return nil
}

func vacancyRouter(vacancy fiber.Router, db *gorm.DB) error {
	vacancy.Get("/:name", func(c *fiber.Ctx) error {
		// var name string
		// obj := new(models.Vacancy)

		// if id = c.Params("name"); name == "" {
		// 	return c.Status(400).JSON(fiber.Map{
		// 		"status":  "error",
		// 		"message": "invalid cpf parameter",
		// 	})
		// }

		return nil
	})

	return nil
}

func professionRouter(profession fiber.Router, db *gorm.DB) error {
	profession.Get("/:name_slug", func(c *fiber.Ctx) error {
		var name_slug string
		obj := new(models.Profession)

		if name_slug = c.Params("name_slug"); name_slug == "" {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": "invalid name slug parameter",
			})
		}

		res := db.Where("name_slug = ?", name_slug).First(obj)
		if err := res.Error; err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}

		if res.RowsAffected <= 0 {
			return c.Status(404).JSON(fiber.Map{
				"status":  "error",
				"message": "record not found",
			})
		}

		data, err := json.Marshal(obj)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}

		m := fiber.Map{}
		if err := json.Unmarshal(data, &m); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"status": "success",
			"data":   m,
		})
	})

	profession.Post("/", func(c *fiber.Ctx) error {
		obj := new(models.Profession)

		if err := c.BodyParser(obj); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}

		res := db.Create(obj)
		if res.Error != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": res.Error.Error(),
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"status": "success",
		})
	})

	profession.Delete("/:name_slug", func(c *fiber.Ctx) error {
		var name_slug string
		obj := new(models.Profession)

		if name_slug = c.Params("name_slug"); name_slug == "" {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": "invalid cpf parameter",
			})
		}
		obj.NameSlug = name_slug

		res := db.Where("name_slug = ?", name_slug).Delete(obj)
		if res.Error != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": res.Error.Error(),
			})
		}

		if res.RowsAffected <= 0 {
			return c.Status(404).JSON(fiber.Map{
				"status":   "error",
				"affected": res.RowsAffected,
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"status":   "success",
			"affected": res.RowsAffected,
		})
	})

	return nil
}

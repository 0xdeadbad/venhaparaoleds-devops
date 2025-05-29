package routes

import (
	"encoding/json"
	"log"

	"github.com/0xdeadbad/venhaparaoleds-devops/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func MainRouter(r fiber.Router, db *gorm.DB) error {
	prof := r.Group("/profession")
	appl := r.Group("/applicant")
	conc := r.Group("/concourse")
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
		c.Accepts("json", "text")
		c.Accepts("application/json")

		var cpf string
		objs := new([]models.Profession)
		obj := new(models.Applicant)

		if cpf = c.Params("cpf"); cpf == "" {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": "invalid cpf parameter",
			})
		}
		obj.CPF = cpf

		db.Model(obj).Association("Professions")
		if err := db.Model(obj).Association("Professions").Find(objs); err != nil {
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

	applicant.Post("/", func(c *fiber.Ctx) error {
		obj := new(models.Applicant)
		// data := c.Body()

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
		obj := new(models.Applicant)

		if cpf = c.Params("cpf"); cpf == "" {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": "invalid cpf parameter",
			})
		}
		obj.CPF = cpf

		res := db.Delete(obj)
		if res.Error != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": res.Error.Error(),
			})
		}

		count := res.Statement.RowsAffected

		log.Printf("}}}}}}}}%d\n", count)

		return c.Status(200).JSON(fiber.Map{
			"status": "success",
		})
	})

	return nil
}

func concourseRouter(concourse fiber.Router, db *gorm.DB) error {
	concourse.Get("/:code", func(c *fiber.Ctx) error {
		c.Accepts("json", "text")
		c.Accepts("application/json")

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
		obj := new(models.Concourse)

		if code = c.Params("code"); code == "" {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": "invalid code parameter",
			})
		}
		obj.ConcCode = code

		res := db.Delete(obj)
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

	concourse.Delete("/:code", func(c *fiber.Ctx) error {
		var code string
		obj := new(models.Concourse)

		if code = c.Params("cpf"); code == "" {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": "invalid cpf parameter",
			})
		}
		obj.ConcCode = code

		res := db.Delete(obj)
		if res.Error != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": res.Error.Error(),
			})
		}

		count := new(int64)
		res.Count(count)

		log.Printf("[[[[[[]]]]]]%d\n", count)

		return c.Status(200).JSON(fiber.Map{
			"status":   "success",
			"affected": count,
		})
	})

	return nil
}

func vacancyRouter(r fiber.Router, db *gorm.DB) error {

	return nil
}

func professionRouter(r fiber.Router, db *gorm.DB) error {

	return nil
}

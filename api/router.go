package api

import "github.com/gofiber/fiber/v2"

type Handler func(*fiber.Ctx) error

// func AddApiGroup[T any](c *fiber.App, g ApiCrud[T]) error {

// 	api := c.Group("/api")
// 	v1 := api.Group("/v1")

// 	return apiV1(v1)
// }

// func apiV1(v1 fiber.Router) error {
// 	err := applicantMethods(v1.Group("/applicant"))

// 	// v1.Group("/concourse", queryConcourseHandle)

// 	return err
// }

// func applicantHandler(c *fiber.Ctx) error {

// 	return nil
// }

// func newApiCrudHandler() (Handler, error) {

// 	return nil, nil
// }

// func applicantMethods(v1 fiber.Router) error {

// 	v1.Get("/:cpf")
// 	v1.Post("/:cpf")
// 	v1.Delete("/:cpf")
// 	v1.Put("/:cpf")

// 	return nil
// }

// func concourseMethods(v1 fiber.Router) error {
// 	v1.Get("/:code")
// 	v1.Post("/:code")
// 	v1.Delete("/:code")
// 	v1.Put("/:code")

// 	return nil
// }

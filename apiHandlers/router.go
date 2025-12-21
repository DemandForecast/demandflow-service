package apiHandlers

import (
	"Suppliers/api"
	"Suppliers/dto"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Router(app *fiber.App, authConfig dto.AuthConfig) {
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	group := app.Group("/Suppliers/api")
	defaultGroup := app.Group("/")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Static("/", "./docs/rapiDoc/build")
	// authMiddleware := NewAuthMiddleware(authConfig)
	// defaultGroup.Use(authMiddleware.ValidateToken)
	DefaultMappings(defaultGroup)
	RouteMappings(group)
}

func RouteMappings(cg fiber.Router) {

cg.Post("/CreateCustomer", api.CreateCustomerApi)
cg.Put("/UpdateCustomer", api.UpdateCustomerApi)
cg.Delete("/DeleteCustomer", api.DeleteCustomerApi)
cg.Get("/FindCustomer", api.FindCustomerApi)
cg.Get("/FindallCustomer", api.FindallCustomerApi)

}

func DefaultMappings(cg fiber.Router) {
	cg.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "Suppliers-APP1924 service is up and running", "version": "1.0"})
	})
}
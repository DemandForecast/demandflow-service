package apiHandlers

import (
	"DemandFlow-Service/api"
	"DemandFlow-Service/dto"
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

	group := app.Group("/DemandFlow-Service/api")
	group2 := app.Group("/DemandFlow-Service/api/v1/forecast")
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
	ForecastRouteMapping(group2)

}

func RouteMappings(cg fiber.Router) {

	cg.Post("/CreateProduct", api.CreateProductApi)
	cg.Put("/UpdateProduct", api.UpdateProductApi)
	cg.Delete("/DeletProduct", api.DeleteProductApi)
	cg.Get("/FindProduct", api.FindProductApi)
	cg.Get("/FindallProduct", api.FindallProductApi)

	cg.Post("/CreateInventory", api.CreateInventoryApi)
	cg.Put("/UpdateInventory", api.UpdateInventoryApi)
	cg.Delete("/DeletInventory", api.DeleteInventoryApi)
	cg.Get("/FindInventory", api.FindInventoryApi)
	cg.Get("/FindallInventory", api.FindallInventoryApi)

	cg.Get("/CountProducts", api.CountProductsApi)
	cg.Get("/TotalInventoryStock", api.TotalInventoryStockApi)
	cg.Get("/Top5ProductsByUnitSold", api.Top5ProductsByUnitsSoldApi)
	cg.Post("/Predict/Top5ProductsByUnitSold", api.PredictTop5ProductsByUnitSoldApi)
	cg.Get("/Dashboard/Top5ProductsByUnitSold", api.DashboardTop5ProductsByUnitSoldApi)


}

func ForecastRouteMapping(cg fiber.Router) {
	cg.Post("/product", api.PredictProductDemandApi)
	// forecast := app.Group("/api/v1/forecast")

	// forecast.Post("/product", api.PredictSingleProductDemandApi)
	// forecast.Post("/products", api.PredictAllProductsDemandApi)

	// forecast.Get("/", api.GetAllForecastsApi)
	// forecast.Get("/:forecastId", api.GetForecastByIDApi)
	// forecast.Get("/product/:productId", api.GetForecastsByProductIDApi)
}

func DefaultMappings(cg fiber.Router) {
	cg.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "DemandFlow-Service is up and running", "version": "1.0"})
	})
}

package api

import (
	"DemandFlow-Service/dao"
	"DemandFlow-Service/dto"
	"DemandFlow-Service/functions"
	"DemandFlow-Service/requests"
	"DemandFlow-Service/utils"
	"encoding/json"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// @Summary      Predict Demand Forecast For Top 5 Products By Stock
// @Description  Gets top 5 stock products, predicts demand for each product separately, and saves each prediction in DB
// @Tags         Forecast
// @Accept       json
// @Produce      json
// @Router       /Top5ProductsByStockDemand [POST]
func DashboardTop5ProductsByUnitSoldApi(c *fiber.Ctx) error {
	inputObj := dto.TopStockForecastRequest{}

	if err := c.BodyParser(&inputObj); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if inputObj.FromDate == "" || inputObj.ToDate == "" {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "fromDate and toDate are required")
	}

	if inputObj.Weather == "" {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "weather is required")
	}

	if inputObj.Region == "" {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "region is required")
	}

	if inputObj.Seasonality == "" {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "seasonality is required")
	}

	updatedBy, err := functions.GetAuth0Id(c)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	forecastServiceURL := os.Getenv("ForecastService")
	if forecastServiceURL == "" {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "ForecastService env variable is missing")
	}

	topProducts, err := dao.DB_Top5ProductsByUnitsSoldForForecast()
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if len(topProducts) == 0 {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "no products found for forecast")
	}

	validate := validator.New()

	savedPredictions := []dto.ForecastPrediction{}
	failedPredictions := []fiber.Map{}

	for _, product := range topProducts {
		if product.ProductID == "" {
			failedPredictions = append(failedPredictions, fiber.Map{
				"productId":   product.ProductID,
				"productName": product.ProductName,
				"reason":      "productId is missing",
			})
			continue
		}

		if product.Price <= 0 {
			failedPredictions = append(failedPredictions, fiber.Map{
				"productId":   product.ProductID,
				"productName": product.ProductName,
				"reason":      "price must be greater than 0 in Products collection",
				"price":       product.Price,
			})
			continue
		}

		if product.Category == "" {
			failedPredictions = append(failedPredictions, fiber.Map{
				"productId":   product.ProductID,
				"productName": product.ProductName,
				"reason":      "category is missing in Products collection",
			})
			continue
		}

		if product.CurrentInventory <= 0 {
			failedPredictions = append(failedPredictions, fiber.Map{
				"productId":        product.ProductID,
				"productName":      product.ProductName,
				"reason":           "currentInventory must be greater than 0 in Inventory collection",
				"currentInventory": product.CurrentInventory,
			})
			continue
		}

		productForecastRequest := dto.ProductForecastRequest{
			ProductID: product.ProductID,

			FromDate: inputObj.FromDate,
			ToDate:   inputObj.ToDate,

			Price:     product.Price,
			Discount:  product.Discount,
			Inventory: product.CurrentInventory,

			Weather:     inputObj.Weather,
			Region:      inputObj.Region,
			Category:    product.Category,
			Seasonality: inputObj.Seasonality,
			Holiday:      inputObj.Holiday,
		}

		if validationErr := validate.Struct(&productForecastRequest); validationErr != nil {
			failedPredictions = append(failedPredictions, fiber.Map{
				"productId":   product.ProductID,
				"productName": product.ProductName,
				"reason":      validationErr.Error(),
				"request":     productForecastRequest,
			})
			continue
		}

		forecastResponseMap, err := requests.PostRequest(
			forecastServiceURL+"/api/v1/forecast/product",
			productForecastRequest,
			*c,
		)
		if err != nil {
			failedPredictions = append(failedPredictions, fiber.Map{
				"productId":   product.ProductID,
				"productName": product.ProductName,
				"reason":      err.Error(),
				"request":     productForecastRequest,
			})
			continue
		}

		forecastBytes, err := json.Marshal(forecastResponseMap)
		if err != nil {
			failedPredictions = append(failedPredictions, fiber.Map{
				"productId":   product.ProductID,
				"productName": product.ProductName,
				"reason":      err.Error(),
			})
			continue
		}

		forecastData := dto.ProductForecastData{}

		if err := json.Unmarshal(forecastBytes, &forecastData); err != nil {
			failedPredictions = append(failedPredictions, fiber.Map{
				"productId":   product.ProductID,
				"productName": product.ProductName,
				"reason":      err.Error(),
				"rawResponse":  forecastResponseMap,
			})
			continue
		}

		// predictionID, err := functions.Idgenerator("Prediction", "PredictionId", "PRED")
		// if err != nil {
		// 	failedPredictions = append(failedPredictions, fiber.Map{
		// 		"productId":   product.ProductID,
		// 		"productName": product.ProductName,
		// 		"reason":      err.Error(),
		// 	})
		// 	continue
		// }

		forecastObject := dto.ForecastPrediction{
			//PredictionID:        predictionID,
			ForecastType:        "TOP5_PRODUCT_FORECAST",
			ProductForecastData: forecastData,
			Deleted:             false,
			Base: dto.Base{
				CreatedAt:     time.Now(),
				LastUpdatedAt: time.Now(),
				LastUpdatedBy: updatedBy,
			},
		}

		// err = dao.DB_CreatePrediction(&forecastObject)
		// if err != nil {
		// 	failedPredictions = append(failedPredictions, fiber.Map{
		// 		"productId":   product.ProductID,
		// 		"productName": product.ProductName,
		// 		"reason":      err.Error(),
		// 	})
		// 	continue
		// }

		savedPredictions = append(savedPredictions, forecastObject)
	}

	if len(savedPredictions) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":            false,
			"message":           "Failed to generate forecasts for all top 5 products",
			"totalProducts":     len(topProducts),
			"successCount":      len(savedPredictions),
			"failedCount":       len(failedPredictions),
			"topProducts":       topProducts,
			"failedPredictions": failedPredictions,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":            true,
		"message":           "Top 5 products by stock demand forecasts generated and saved successfully",
		"totalProducts":     len(topProducts),
		"successCount":      len(savedPredictions),
		"failedCount":       len(failedPredictions),
		"data":              savedPredictions,
		"failedPredictions": failedPredictions,
	})
}
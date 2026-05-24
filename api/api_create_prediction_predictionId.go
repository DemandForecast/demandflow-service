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

func PredictProductDemandApi(c *fiber.Ctx) error {
	inputObj := dto.ProductForecastRequest{}

	if err := c.BodyParser(&inputObj); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if inputObj.ProductID == "" {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "productId is required")
	}

	if inputObj.FromDate == "" || inputObj.ToDate == "" {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "fromDate and toDate are required")
	}

	validate := validator.New()
	if validationErr := validate.Struct(&inputObj); validationErr != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
	}

	updatedBy, err := functions.GetAuth0Id(c)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	forecastResponseMap, err := requests.PostRequest(
		os.Getenv("ForecastService")+"/api/v1/forecast/product",
		inputObj,
		*c,
	)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	forecastBytes, err := json.Marshal(forecastResponseMap)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	forecastData := dto.ProductForecastData{}

	if err := json.Unmarshal(forecastBytes, &forecastData); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	predictionID, err := functions.Idgenerator("Prediction", "PredictionId", "PRED")
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	forecastObject := dto.ForecastPrediction{
		PredictionID:        predictionID,
		ForecastType:        "PRODUCT_FORECAST",
		ProductForecastData: forecastData,
		Deleted:             false,
		Base: dto.Base{
			CreatedAt:     time.Now(),
			LastUpdatedAt: time.Now(),
			LastUpdatedBy: updatedBy,
		},
	}

	err = dao.DB_CreatePrediction(&forecastObject)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":       true,
		"message":      "Product demand forecast generated and saved successfully",
		"predictionId": predictionID,
		"data":         forecastObject,
	})
}
package api

import (
	"DemandFlow-Service/utils"
	"time"

	"github.com/gofiber/fiber/v2"

	"DemandFlow-Service/functions"

	"DemandFlow-Service/dto"

	"github.com/go-playground/validator/v10"

	"DemandFlow-Service/dao"
)

// @Summary      CreateProduct
// @Description   This API performs the POST operation on DemandFlow-Service. It allows you to create DemandFlow-Service records.
// @Tags          DemandFlow-Service
// @Accept       json
// @Produce      json
// @Param        data body dto.DemandFlow-Service false "string collection"
// @Success      200  {array}   dto.DemandFlow-Service "Status OK"
// @Success      202  {array}   dto.DemandFlow-Service "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /CreateProduct [POST]

func CreateProductApi(c *fiber.Ctx) error {

	inputObj := dto.Product{}

	if err := c.BodyParser(&inputObj); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	CustomerId, err := functions.Idgenerator("Products", "ProductId", "Pro")
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	inputObj.ProductId = CustomerId

	validate := validator.New()
	if validationErr := validate.Struct(&inputObj); validationErr != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
	}

	updatedBy, err := functions.GetAuth0Id(c)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	inputObj.Base = dto.Base{
		CreatedAt:     time.Now(),
		LastUpdatedAt: time.Now(),
		LastUpdatedBy: updatedBy,
	}

	err = dao.DB_CreateProduct(&inputObj)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SendSuccessResponse(c)

}

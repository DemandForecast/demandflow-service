package api

import (
	"DemandFlow-Service/functions"
	"DemandFlow-Service/utils"
	"time"

	"github.com/gofiber/fiber/v2"

	"DemandFlow-Service/dto"

	"github.com/go-playground/validator/v10"

	"DemandFlow-Service/dao"
)

// @Summary      UpdateSupplier
// @Description   This API performs the PUT operation on DemandFlow-Service. It allows you to update DemandFlow-Service records.
// @Tags          DemandFlow-Service
// @Accept       json
// @Produce      json
// @Param        data body dto.DemandFlow-Service false "string collection"
// @Success      200  {array}   dto.DemandFlow-Service "Status OK"
// @Success      202  {array}   dto.DemandFlow-Service "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateSupplier [PUT]

func UpdateProductApi(c *fiber.Ctx) error {

	inputObj := dto.Product{}

	if err := c.BodyParser(&inputObj); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	// organizationId, err := functions.GetOrganizationId(c)
	// if err != nil {
	// 	return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	// }

	// inputObj.OrganizationId = organizationId

	validate := validator.New()
	if validationErr := validate.Struct(&inputObj); validationErr != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
	}

	updatedBy, err := functions.GetAuth0Id(c)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	inputObj.Base = dto.Base{
		LastUpdatedAt: time.Now(),
		LastUpdatedBy: updatedBy,
	}

	err = dao.DB_UpdateProduct(&inputObj)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SendSuccessResponse(c)

}

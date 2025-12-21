package api

import (
	"Suppliers/functions"
	"Suppliers/utils"
	"time"

	"github.com/gofiber/fiber/v2"

	"Suppliers/dto"

	"github.com/go-playground/validator/v10"

	"Suppliers/dao"
)

// @Summary      UpdateSupplier
// @Description   This API performs the PUT operation on Supplier. It allows you to update Supplier records.
// @Tags          Supplier
// @Accept       json
// @Produce      json
// @Param        data body dto.Supplier false "string collection"
// @Success      200  {array}   dto.Supplier "Status OK"
// @Success      202  {array}   dto.Supplier "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateSupplier [PUT]

func UpdateCustomerApi(c *fiber.Ctx) error {

	inputObj := dto.Customer{}

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

	err = dao.DB_UpdateCustomer(&inputObj)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SendSuccessResponse(c)

}

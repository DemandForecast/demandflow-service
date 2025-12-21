package api

import (
	"Suppliers/utils"
	"time"

	"github.com/gofiber/fiber/v2"

	"Suppliers/functions"

	"Suppliers/dto"

	"github.com/go-playground/validator/v10"

	"Suppliers/dao"
)

// @Summary      CreateSupplier
// @Description   This API performs the POST operation on Supplier. It allows you to create Supplier records.
// @Tags          Supplier
// @Accept       json
// @Produce      json
// @Param        data body dto.Supplier false "string collection"
// @Success      200  {array}   dto.Supplier "Status OK"
// @Success      202  {array}   dto.Supplier "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /CreateSupplier [POST]

func CreateCustomerApi(c *fiber.Ctx) error {

	inputObj := dto.Customer{}

	if err := c.BodyParser(&inputObj); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	// organizationId, err := functions.GetOrganizationId(c)
    // if err != nil {
	// 	return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	// }

	// inputObj.OrganizationId = organizationId

	CustomerId, err := functions.Idgenerator("Customers", "CustomerId", "Cus")
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	inputObj.CustomerId = CustomerId

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

	err = dao.DB_CreateCustomer(&inputObj)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SendSuccessResponse(c)

}

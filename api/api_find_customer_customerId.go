package api

import (
	"Suppliers/utils"

	"github.com/gofiber/fiber/v2"

	"Suppliers/dao"
)

// @Summary      FindSupplier
// @Description   This API performs the GET operation on Supplier. It allows you to retrieve Supplier records.
// @Tags          Supplier
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Supplier "Status OK"
// @Success      202  {array}   dto.Supplier "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindSupplier [GET]

func FindCustomerApi(c *fiber.Ctx) error {

	CustomerId := c.Query("CustomerId")

    // organizationId, err := functions.GetOrganizationId(c)
	// if err != nil {
	// 	return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	// }

	returnValue, err := dao.DB_FindCustomerbyCustomerId(CustomerId)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}

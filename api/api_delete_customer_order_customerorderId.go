package api

import (
	"Suppliers/functions"
	"Suppliers/utils"

	"github.com/gofiber/fiber/v2"

	"Suppliers/dao"
)

// @Summary      DeleteSupplier
// @Description   This API performs the DELETE operation on Supplier. It allows you to delete Supplier records.
// @Tags          Supplier
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Supplier "Status OK"
// @Success      202  {array}   dto.Supplier "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteSupplier [DELETE]

func DeleteCustomerOrderApi(c *fiber.Ctx) error {

	CustomerOrderId := c.Query("CustomerOrderId")

    organizationId, err := functions.GetOrganizationId(c)
	organizationId = organizationId
	
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	err = dao.DB_DeleteCustomerOrder(CustomerOrderId)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SendSuccessResponse(c)

}

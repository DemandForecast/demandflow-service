package api

import (
	"DemandFlow-Service/functions"
	"DemandFlow-Service/utils"

	"github.com/gofiber/fiber/v2"

	"DemandFlow-Service/dao"
)

// @Summary      DeleteProduct
// @Description   This API performs the DELETE operation on DemandFlow-Service. It allows you to delete DemandFlow-Service records.
// @Tags          DemandFlow-Service
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.DemandFlow-Service "Status OK"
// @Success      202  {array}   dto.DemandFlow-Service "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteSupplier [DELETE]

func DeleteProductApi(c *fiber.Ctx) error {

	ProductId := c.Query("ProductId")

	organizationId, err := functions.GetOrganizationId(c)
	organizationId = organizationId

	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	err = dao.DB_DeleteProduct(ProductId)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SendSuccessResponse(c)

}

package api

import (
	"DemandFlow-Service/utils"

	"github.com/gofiber/fiber/v2"

	"DemandFlow-Service/dao"
)

// @Summary      FindSupplier
// @Description   This API performs the GET operation on DemandFlow-Service. It allows you to retrieve DemandFlow-Service records.
// @Tags          DemandFlow-Service
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.DemandFlow-Service "Status OK"
// @Success      202  {array}   dto.DemandFlow-Service "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindSupplier [GET]

func FindProductApi(c *fiber.Ctx) error {

	ProductId := c.Query("ProductId")

	// organizationId, err := functions.GetOrganizationId(c)
	// if err != nil {
	// 	return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	// }

	returnValue, err := dao.DB_FindProductbyProductId(ProductId)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}

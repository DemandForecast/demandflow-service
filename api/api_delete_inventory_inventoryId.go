package api

import (
	"DemandFlow-Service/functions"
	"DemandFlow-Service/utils"

	"github.com/gofiber/fiber/v2"

	"DemandFlow-Service/dao"
)

// @Summary      DeleteInventory
// @Description   This API performs the DELETE operation on DemandFlow-Service. It allows you to delete DemandFlow-Service records.
// @Tags          DemandFlow-Service
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.DemandFlow-Service "Status OK"
// @Success      202  {array}   dto.DemandFlow-Service "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteInventory [DELETE]

func DeleteInventoryApi(c *fiber.Ctx) error {

	InventoryId := c.Query("InventoryId")

	organizationId, err := functions.GetOrganizationId(c)
	organizationId = organizationId

	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	err = dao.DB_DeleteInventory(InventoryId)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SendSuccessResponse(c)

}

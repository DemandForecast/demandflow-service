package api

import (
	"DemandFlow-Service/utils"

	"github.com/gofiber/fiber/v2"

	"DemandFlow-Service/dao"
)

// @Summary      FindallProduct
// @Description   This API performs the GET operation on DemandFlow-Service. It allows you to retrieve DemandFlow-Service records.
// @Tags          DemandFlow-Service
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.DemandFlow-Service "Status OK"
// @Success      202  {array}   dto.DemandFlow-Service "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindallProduct [GET]

func FindallProductsWithPgApi(c *fiber.Ctx) error {

	page := c.Query("page", "1")
	size := c.Query("size", "10")
	search := c.Query("searchkeyword", "")

	count, Products, err := dao.DB_FindProductsWithPg(page, size, search)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	resp := map[string]interface{}{
		"count":    count,
		"Products": Products,
	}

	return c.Status(fiber.StatusAccepted).JSON(resp)

}

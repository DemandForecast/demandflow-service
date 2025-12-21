package api

import (
	"Suppliers/utils"

	"github.com/gofiber/fiber/v2"

	"Suppliers/dao"
)

// @Summary      FindallSupplier
// @Description   This API performs the GET operation on Supplier. It allows you to retrieve Supplier records.
// @Tags          Supplier
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Supplier "Status OK"
// @Success      202  {array}   dto.Supplier "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindallSupplier [GET]

func FindallCustomersWithPgApi(c *fiber.Ctx) error {


	page := c.Query("page", "1")
    size := c.Query("size", "10")
    search := c.Query("searchkeyword", "")


    count, customers, err := dao.DB_FindCustomersWithPg( page, size, search)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }

    resp := map[string]interface{}{
        "count":    count,
        "customers": customers,
    }

    return c.Status(fiber.StatusAccepted).JSON(resp)

}

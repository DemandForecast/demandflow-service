package api

import (
	"DemandFlow-Service/dao"
	"DemandFlow-Service/utils"

	"github.com/gofiber/fiber/v2"
)

// @Summary      Total Inventory Stock
// @Description  Returns total stock quantity
// @Tags         DemandFlow-Service
// @Accept       json
// @Produce      json
// @Router       /TotalInventoryStock [GET]

func TotalInventoryStockApi(c *fiber.Ctx) error {

	totalStock, err := dao.DB_TotalInventoryStock()
	if err != nil {
		return utils.SendErrorResponse(
			c,
			fiber.StatusBadRequest,
			err.Error(),
		)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"totalStock": totalStock,
	})
}
package api

import (
	"DemandFlow-Service/dao"
	"DemandFlow-Service/utils"

	"github.com/gofiber/fiber/v2"
)

// @Summary      Top 5 Products by Stock
// @Description  Returns top 5 products with highest stock
// @Tags         Dashboard
// @Accept       json
// @Produce      json
// @Router       /Top5ProductsByStock [GET]

func Top5ProductsByStockApi(c *fiber.Ctx) error {

	result, err := dao.DB_Top5ProductsByStock()
	if err != nil {
		return utils.SendErrorResponse(
			c,
			fiber.StatusBadRequest,
			err.Error(),
		)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    result,
	})
}
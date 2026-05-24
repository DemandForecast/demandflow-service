package api

import (
	"DemandFlow-Service/dao"
	"DemandFlow-Service/utils"

	"github.com/gofiber/fiber/v2"
)

// @Summary      Top 5 Products by Units Sold
// @Description  Returns top 5 products with highest units sold
// @Tags         Dashboard
// @Accept       json
// @Produce      json
// @Router       /Top5ProductsByUnitsSold [GET]

func Top5ProductsByUnitsSoldApi(c *fiber.Ctx) error {

	result, err := dao.DB_Top5ProductsByUnitsSold()
	if err != nil {
		return utils.SendErrorResponse(
			c,
			fiber.StatusBadRequest,
			err.Error(),
		)
	}

	// return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 	"data": result,
	// })
	return c.Status(fiber.StatusAccepted).JSON(&result)
}
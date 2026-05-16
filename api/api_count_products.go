package api

import (
	"DemandFlow-Service/dao"
	"DemandFlow-Service/utils"

	"github.com/gofiber/fiber/v2"
)

// @Summary      Count Products
// @Description  Returns total active products count
// @Tags         DemandFlow-Service
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Router       /CountProducts [GET]

func CountProductsApi(c *fiber.Ctx) error {

	count, err := dao.DB_CountProducts()
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		// "success": true,
		"count":   count,
	})
}
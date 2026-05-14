package functions

import (
	"DemandFlow-Service/dto"

	"github.com/gofiber/fiber/v2"
)

func CheckUserAccess(c *fiber.Ctx, organizationId string) error {
	user, ok := c.Locals("user").(dto.UserCacheData)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized access",
		})
	}

	if user.OrganizationId != organizationId {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "unauthorized access",
		})
	}

	return nil
}

func GetOrganizationId(c *fiber.Ctx) (string, error) {
	user, ok := c.Locals("user").(dto.UserCacheData)
	if !ok {
		return "", c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized access",
		})
	}

	if user.OrganizationId == "" {
		return "", c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized access",
		})
	}

	return user.OrganizationId, nil
}

func GetAuth0Id(c *fiber.Ctx) (string, error) {
	user, ok := c.Locals("user").(dto.UserCacheData)
	if !ok {
		return "", c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized access",
		})
	}

	if user.Sub == "" {
		return "", c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized access",
		})
	}

	return user.Sub, nil
}

package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func RequireAuth(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
	}

	payload, err := Decoder(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	c.Locals("user", payload.User)
	c.Locals("role", payload.Role)

	return c.Next()
}

func Middlewares(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authToken := c.Get("Authorization")
		if authToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization Header"})
		}
		dataDecode, err := Decoder(authToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
		}
		if dataDecode.Role != requiredRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "This User not Permitted for this function"})
		}

		return c.Next()
	}
}

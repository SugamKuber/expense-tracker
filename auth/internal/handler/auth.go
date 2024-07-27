package handler

import (
    "auth/lib/model"
    "auth/util"
    "github.com/gofiber/fiber/v2"
    "net/http"
)

func AuthMiddleware(c *fiber.Ctx) error {
    token := c.Get("Authorization")
    if token == "" {
        return c.Status(http.StatusUnauthorized).SendString("Missing token")
    }

    userID, err := util.ParseJWT(token)
    if err != nil {
        return c.Status(http.StatusUnauthorized).SendString("Invalid token")
    }

    user, err := model.GetUserByID(userID)
    if err != nil {
        return c.Status(http.StatusUnauthorized).SendString("User not found")
    }

    c.Locals("user", user)
    return c.Next()
}

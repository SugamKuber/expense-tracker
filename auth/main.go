package main

import (
    "github.com/gofiber/fiber/v2"
)

type User struct {
    Email    string `json:"email"`
    Name     string `json:"name"`
    Mobile   string `json:"mobile"`
}

var users = map[string]User{}

func main() {
    app := fiber.New()

    app.Post("/user", func(c *fiber.Ctx) error {
        user := new(User)
        if err := c.BodyParser(user); err != nil {
            return c.Status(400).SendString(err.Error())
        }
        users[user.Email] = *user
        return c.Status(201).JSON(user)
    })

    app.Get("/user/:email", func(c *fiber.Ctx) error {
        email := c.Params("email")
        user, exists := users[email]
        if !exists {
            return c.Status(404).SendString("User not found")
        }
        return c.JSON(user)
    })

    app.Listen(":3000")
}

package main

import (
    "github.com/gofiber/fiber/v2"
)

type Expense struct {
    ID          int             `json:"id"`
    Amount      float64         `json:"amount"`
    SplitType   string          `json:"split_type"`
    Participants map[string]float64 `json:"participants"`
}

var expenses = map[int]Expense{}
var nextID = 1

func main() {
    app := fiber.New()

    app.Post("/expense", func(c *fiber.Ctx) error {
        expense := new(Expense)
        if err := c.BodyParser(expense); err != nil {
            return c.Status(400).SendString(err.Error())
        }
        expense.ID = nextID
        nextID++
        expenses[expense.ID] = *expense
        return c.Status(201).JSON(expense)
    })

    app.Get("/expense/:id", func(c *fiber.Ctx) error {
        id := c.Params("id")
        expense, exists := expenses[id]
        if !exists {
            return c.Status(404).SendString("Expense not found")
        }
        return c.JSON(expense)
    })

    app.Listen(":3001")
}

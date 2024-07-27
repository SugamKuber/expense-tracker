package main

import (
    "github.com/gofiber/fiber/v2"
    "encoding/csv"
    "os"
)

type Balance struct {
    User     string  `json:"user"`
    Amount   float64 `json:"amount"`
}

var balances = map[string]Balance{}

func main() {
    app := fiber.New()

    app.Get("/balances", func(c *fiber.Ctx) error {
        return c.JSON(balances)
    })

    app.Get("/balance_sheet/download", func(c *fiber.Ctx) error {
        file, err := os.Create("balance_sheet.csv")
        if err != nil {
            return c.Status(500).SendString(err.Error())
        }
        defer file.Close()

        writer := csv.NewWriter(file)
        defer writer.Flush()

        for _, balance := range balances {
            record := []string{balance.User, fmt.Sprintf("%f", balance.Amount)}
            if err := writer.Write(record); err != nil {
                return c.Status(500).SendString(err.Error())
            }
        }

        return c.SendFile("balance_sheet.csv")
    })

    app.Listen(":3002")
}

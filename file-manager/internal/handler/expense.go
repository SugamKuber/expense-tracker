package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
	"file-manager/lib/model"
)

func handleTrack(c *fiber.Ctx, endpoint string) error {
	user, ok := c.Locals("user").(*model.User)
	if !ok {
		return c.Status(http.StatusInternalServerError).SendString("User not found")
	}

	token := c.Get("Authorization")

	client := &http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to create request")
	}
	req.Header.Set("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to make request")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to read response")
	}

	fmt.Println(user)

	return c.Status(resp.StatusCode).Send(body)
}

func generateExcelForTrackAll(c *fiber.Ctx, endpoint string) error {
	err := handleTrack(c, endpoint)
	if err != nil {
		return err
	}

	body := c.Response().Body()

	var expensesResponse ExpensesResponse
	if err := json.Unmarshal(body, &expensesResponse); err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to parse response")
	}

	buffer := new(bytes.Buffer)
	if err := createExcelForTrackAll(expensesResponse.Expenses, buffer); err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to create Excel file")
	}

	c.Response().Header.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header.Set("Content-Disposition", `attachment; filename="trackall_expenses.xlsx"`)
	return c.SendStream(buffer)
}

func generateExcelForTrackMe(c *fiber.Ctx, endpoint string) error {
	err := handleTrack(c, endpoint)
	if err != nil {
		return err
	}

	body := c.Response().Body()

	var myExpenses MyExpenses
	if err := json.Unmarshal(body, &myExpenses); err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to parse response")
	}

	buffer := new(bytes.Buffer)
	if err := createExcelForTrackMe(myExpenses.Expenses, buffer); err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to create Excel file")
	}

	c.Response().Header.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header.Set("Content-Disposition", `attachment; filename="trackme_expenses.xlsx"`)
	return c.SendStream(buffer)
}

func TrackMe(c *fiber.Ctx) error {
	return generateExcelForTrackMe(c, "http://localhost:3001/track/me")
}

func TrackAll(c *fiber.Ctx) error {
	return generateExcelForTrackAll(c, "http://localhost:3001/track/all")
}
